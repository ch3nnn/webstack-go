/**
 * @Author: chentong
 * @Date: 2025/02/12 14:25
 */

package config

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/duke-git/lancet/v2/pointer"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/middleware"
	configservice "github.com/ch3nnn/webstack-go/internal/service/config"
	"github.com/ch3nnn/webstack-go/pkg/config"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

var (
	logger *log.Logger
	hdl    *handler.Handler
	j      *jwt.JWT
	router *gin.Engine
)

func TestMain(m *testing.M) {
	err := os.Setenv("APP_CONF", "../../../config/test.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}

	envConf := flag.String("conf", "config/test.yml", "config path, eg: -conf ./config/test.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logPath := filepath.Join("../../../", conf.GetString("log.log_file_name"))
	conf.Set("log.log_file_name", logPath)

	logger = log.NewLog(conf)
	hdl = handler.NewHandler(logger)
	j = jwt.NewJwt(conf)

	gin.SetMode(gin.TestMode)
	router = gin.Default()
	router.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
	)

	code := m.Run()
	os.Exit(code)
}

func TestConfigHandler_Config(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfigService := configservice.NewMockService(ctrl)
	mockConfigService.EXPECT().GetConfig(gomock.Any()).Return(&v1.ConfigResp{
		SiteTitle:   "测试站点",
		SiteKeyword: "测试关键词",
		SiteDesc:    "测试描述",
		AboutSite:   "关于站点",
		AboutAuthor: "关于作者",
		IsAbout:     true,
	}, nil)

	configHandler := NewHandler(hdl, mockConfigService)
	router.GET("/config", configHandler.Config)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/config", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp v1.Response
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

func TestConfigHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updateReq := &v1.ConfigUpdateReq{
		SiteTitle:   pointer.Of("新站点标题"),
		SiteKeyword: pointer.Of("新关键词"),
		SiteDesc:    pointer.Of("新描述"),
		AboutSite:   pointer.Of("新站点介绍"),
		AboutAuthor: pointer.Of("新作者介绍"),
		IsAbout:     pointer.Of(true),
	}

	mockConfigService := configservice.NewMockService(ctrl)
	mockConfigService.EXPECT().Update(gomock.Any(), updateReq).Return(&v1.ConfigUpdateResp{}, nil)

	configHandler := NewHandler(hdl, mockConfigService)
	router.PUT("/config", configHandler.Update)

	body, _ := json.Marshal(updateReq)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/config", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
