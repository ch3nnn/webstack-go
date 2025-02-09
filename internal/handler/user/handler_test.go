/**
 * @Author: chentong
 * @Date: 2025/02/08 22:05
 */

package user

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	httpexpect "github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/middleware"
	userservice "github.com/ch3nnn/webstack-go/internal/service/user"
	"github.com/ch3nnn/webstack-go/pkg/config"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

var userId = 1

var logger *log.Logger

var hdl *handler.Handler

var j *jwt.JWT

var router *gin.Engine

func performRequest(r http.Handler, method, path string, body *bytes.Buffer) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}

func genToken(t *testing.T) string {
	token, err := j.GenToken(userId, time.Now().Add(time.Hour*24))
	if err != nil {
		t.Error(err)
		return token
	}
	return token
}

func newHttpExcept(t *testing.T, router *gin.Engine) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(router),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: nil,
	})
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	err := os.Setenv("APP_CONF", "../../../config/test.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}

	envConf := flag.String("conf", "config/test.yml", "config path, eg: -conf ./config/test.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	// modify log directory
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

func TestUserHandler_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := userservice.NewMockService(ctrl)
	userHandler := NewHandler(hdl, mockUserService)
	router.GET("/logout", userHandler.Logout)

	obj := newHttpExcept(t, router).GET("/logout").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.IsEmpty()
}

func TestUserHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	params := v1.LoginReq{
		Username: "admin",
		Password: "admin",
	}

	tk := genToken(t)
	mockUserService := userservice.NewMockService(ctrl)
	mockUserService.EXPECT().Login(gomock.Any(), &params).Return(&v1.LoginResp{Token: tk}, nil)

	userHandler := NewHandler(hdl, mockUserService)
	router.POST("/login", userHandler.Login)

	obj := newHttpExcept(t, router).POST("/login").
		WithHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		WithForm(params).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("token").IsEqual(tk)
}

func TestUserHandler_Info(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "admin"
	mockUserService := userservice.NewMockService(ctrl)
	mockUserService.EXPECT().Info(gomock.Any(), nil).Return(&v1.InfoResp{
		Username: username,
		Menus:    nil,
	}, nil)

	userHandler := NewHandler(hdl, mockUserService)
	router.Use(middleware.StrictAuth(j, logger))
	router.GET("/info", userHandler.Info)

	obj := newHttpExcept(t, router).GET("/info").
		WithHeader("Token", genToken(t)).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.Value("username").IsEqual(username)
}

func TestUserHandler_UpdatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &v1.UpdatePasswordReq{
		OldPassword: "admin",
		NewPassword: "admin",
	}
	mockUserService := userservice.NewMockService(ctrl)
	mockUserService.EXPECT().UpdatePassword(gomock.Any(), req).Return(&v1.UpdatePasswordResp{}, nil)

	userHandler := NewHandler(hdl, mockUserService)
	router.Use(middleware.StrictAuth(j, logger))
	router.PUT("/update_password", userHandler.UpdatePassword)

	obj := newHttpExcept(t, router).PUT("/update_password").
		WithHeader("Token", genToken(t)).
		WithHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		WithForm(req).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()

	obj.IsEmpty()
}
