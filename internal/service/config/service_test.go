/**
 * @Author: chentong
 * @Date: 2025/02/11 13:34
 */

package config

import (
	"flag"
	"mime/multipart"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/duke-git/lancet/v2/pointer"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
	"github.com/ch3nnn/webstack-go/pkg/config"
	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

var (
	logger *log.Logger
	j      *jwt.JWT
)

func setupRepository(t *testing.T) (*repository.Repository, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm connection: %v", err)
	}

	query.SetDefault(db)

	return repository.NewRepository(logger, db), mock
}

func TestMain(m *testing.M) {
	err := os.Setenv("APP_CONF", "../../../config/test.yml")
	if err != nil {
		panic(err)
	}

	envConf := flag.String("conf", "config/test.yml", "config path, eg: -conf ./config/test.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	j = jwt.NewJwt(conf)

	code := m.Run()
	os.Exit(code)
}

func TestService_GetConfig(t *testing.T) {
	sysConfig := &model.SysConfig{
		ID:          1,
		AboutSite:   "config",
		AboutAuthor: "config",
		IsAbout:     true,
		SiteTitle:   "config",
		SiteKeyword: "config",
		SiteDesc:    "config",
		SiteRecord:  "config",
		SiteLogo:    "config",
		SiteFavicon: "config",
	}

	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	configDao := repository.NewMockISysConfigDao(ctrl)
	configFunc := repository.NewMockiCustomGenSysConfigFunc(ctrl)

	configDao.EXPECT().WithContext(gomock.Any()).Return(configFunc)
	configFunc.EXPECT().FindOne(gomock.Any()).Return(sysConfig, nil)

	configService := NewService(srv, configDao)
	_, err := configService.GetConfig(ctx)

	assert.NoError(t, err)
}

func TestService_Update(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	configDao := repository.NewMockISysConfigDao(ctrl)
	configFunc := repository.NewMockiCustomGenSysConfigFunc(ctrl)

	configDao.EXPECT().WithContext(gomock.Any()).Return(configFunc)
	configDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	configFunc.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	configService := NewService(srv, configDao)
	_, err := configService.Update(ctx, &v1.ConfigUpdateReq{
		AboutSite:   pointer.Of("config"),
		AboutAuthor: pointer.Of("config"),
		IsAbout:     pointer.Of(true),
		SiteTitle:   pointer.Of("config"),
		SiteKeyword: pointer.Of("config"),
		SiteDesc:    pointer.Of("config"),
		SiteRecord:  pointer.Of("config"),
		LogFile:     &multipart.FileHeader{},
		FaviconFile: &multipart.FileHeader{},
	})

	assert.NoError(t, err)
}
