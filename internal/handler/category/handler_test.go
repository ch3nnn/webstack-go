/**
 * @Author: chentong
 * @Date: 2025/02/12 13:25
 */

package category

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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/middleware"
	categoryservice "github.com/ch3nnn/webstack-go/internal/service/category"
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
	fmt.Println("begin")
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
	fmt.Println("test end")
	os.Exit(code)
}

func TestCategoryHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryService := categoryservice.NewMockService(ctrl)
	mockCategoryService.EXPECT().List(gomock.Any(), nil).Return(&v1.CategoryListResp{
		List: []v1.CategoryList{},
	}, nil)

	categoryHandler := NewHandler(hdl, mockCategoryService)
	router.GET("/categories", categoryHandler.List)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/categories", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoryHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &v1.CategoryCreateReq{
		Name: "测试分类",
		Icon: "test-icon",
	}

	mockCategoryService := categoryservice.NewMockService(ctrl)
	mockCategoryService.EXPECT().Create(gomock.Any(), req).Return(&v1.CategoryCreateResp{
		Category: v1.Category{
			ID:        1,
			ParentID:  0,
			Sort:      0,
			Title:     "测试分类",
			Icon:      "test-icon",
			IsUsed:    true,
			Level:     1,
			CreatedAt: nil,
			UpdatedAt: nil,
		},
	}, nil)

	categoryHandler := NewHandler(hdl, mockCategoryService)
	router.POST("/categories", categoryHandler.Create)

	w := httptest.NewRecorder()
	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", "/categories", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoryHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &v1.CategoryDeleteReq{
		ID: 1,
	}

	mockCategoryService := categoryservice.NewMockService(ctrl)
	mockCategoryService.EXPECT().Delete(gomock.Any(), req).Return(&v1.CategoryDeleteResp{}, nil)

	categoryHandler := NewHandler(hdl, mockCategoryService)
	router.DELETE("/categories/:id", categoryHandler.Delete)

	w := httptest.NewRecorder()
	httpReq, _ := http.NewRequest("DELETE", "/categories/1", nil)
	router.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)
}
