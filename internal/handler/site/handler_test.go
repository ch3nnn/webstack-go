/**
 * @Author: chentong
 * @Date: 2025/02/12 16:45
 */

package site

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
	"github.com/xuri/excelize/v2"
	"go.uber.org/mock/gomock"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/middleware"
	siteservice "github.com/ch3nnn/webstack-go/internal/service/site"
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

func TestSiteHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 构造请求参数
	siteListReq := &v1.SiteListReq{
		Page:       1,
		PageSize:   10,
		Search:     "test",
		CategoryID: 1,
	}

	// Mock service 返回
	mockSiteService := siteservice.NewMockService(ctrl)
	mockSiteService.EXPECT().List(gomock.Any(), siteListReq).Return(&v1.SiteListResp{
		List: []v1.Site{
			{
				Id:          1,
				Title:       "测试站点",
				Url:         "http://test.com",
				CategoryId:  1,
				Category:    "测试分类",
				Description: "测试描述",
				IsUsed:      true,
				Sort:        1,
			},
		},
		Pagination: v1.SiteLisPagination{
			Total:        1,
			CurrentPage:  1,
			PerPageCount: 10,
		},
	}, nil)

	siteHandler := NewHandler(hdl, mockSiteService)
	router.GET("/sites", siteHandler.List)

	// 构造请求
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/sites?page=1&page_size=10&search=test&category_id=1", nil)
	router.ServeHTTP(w, req)

	// 验证结果
	assert.Equal(t, http.StatusOK, w.Code)

	var resp v1.Response
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

func TestSiteHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 构造请求参数
	req := &v1.SiteCreateReq{
		Url:        "http://test.com",
		CategoryID: 1,
		IsUsed:     true,
	}

	// Mock service 返回
	mockSiteService := siteservice.NewMockService(ctrl)
	mockSiteService.EXPECT().BatchCreate(gomock.Any(), req).Return(&v1.SiteCreateResp{
		SuccessCount: 1,
		FailCount:    0,
	}, nil)

	siteHandler := NewHandler(hdl, mockSiteService)
	router.POST("/sites", siteHandler.Create)

	// 构造请求
	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	httpReq, _ := http.NewRequest("POST", "/sites", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, httpReq)

	// 验证结果
	assert.Equal(t, http.StatusOK, w.Code)

	var resp v1.Response
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

func TestSiteHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteID := 1

	// Mock service 返回
	mockSiteService := siteservice.NewMockService(ctrl)
	mockSiteService.EXPECT().Delete(gomock.Any(), &v1.SiteDeleteReq{ID: siteID}).Return(&v1.SiteDeleteResp{
		ID: siteID,
	}, nil)

	siteHandler := NewHandler(hdl, mockSiteService)
	router.DELETE("/sites/:id", siteHandler.Delete)

	// 构造请求
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/sites/%d", siteID), nil)
	router.ServeHTTP(w, req)

	// 验证结果
	assert.Equal(t, http.StatusOK, w.Code)

	var resp v1.Response
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

func TestSiteHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 构造请求参数
	req := &v1.SiteUpdateReq{
		Id:          1,
		Title:       "更新站点",
		Url:         "http://update.com",
		CategoryId:  2,
		Icon:        "update-icon",
		Description: "更新描述",
		IsUsed:      pointer.Of(true),
		Sort:        2,
	}

	// Mock service 返回
	mockSiteService := siteservice.NewMockService(ctrl)
	mockSiteService.EXPECT().Update(gomock.Any(), req).Return(&v1.SiteUpdateResp{
		ID: req.Id,
	}, nil)

	siteHandler := NewHandler(hdl, mockSiteService)
	router.PUT("/sites/:id", siteHandler.Update)

	// 构造请求
	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	httpReq, _ := http.NewRequest("PUT", fmt.Sprintf("/sites/%d", req.Id), bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, httpReq)

	// 验证结果
	assert.Equal(t, http.StatusOK, w.Code)

	var resp v1.Response
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Code)
}

func TestSiteHandler_Export(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 构造请求参数
	req := &v1.SiteExportReq{
		Search:     "test",
		CategoryID: 1,
	}

	// Mock service 返回
	mockSiteService := siteservice.NewMockService(ctrl)
	mockSiteService.EXPECT().Export(gomock.Any(), req).Return(&v1.SiteExportResp{
		File: excelize.NewFile(),
	}, nil)

	siteHandler := NewHandler(hdl, mockSiteService)
	router.GET("/sites/export", siteHandler.Export)

	// 构造请求
	w := httptest.NewRecorder()
	httpReq, _ := http.NewRequest("GET", "/sites/export?search=test&category_id=1", nil)
	router.ServeHTTP(w, httpReq)

	// 验证结果
	assert.Equal(t, http.StatusOK, w.Code)
}
