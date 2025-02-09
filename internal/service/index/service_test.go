/**
 * @Author: chentong
 * @Date: 2025/02/09 21:45
 */

package index

import (
	"flag"
	"os"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

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
	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		panic(err)
	}

	envConf := flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	j = jwt.NewJwt(conf)

	code := m.Run()
	os.Exit(code)
}

func TestService_About(t *testing.T) {
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

	siteDao := repository.NewMockIStSiteDao(ctrl)

	categoryDao := repository.NewMockIStCategoryDao(ctrl)

	configDao := repository.NewMockISysConfigDao(ctrl)
	configFunc := repository.NewMockiCustomGenSysConfigFunc(ctrl)

	configDao.EXPECT().WithContext(gomock.Any()).Return(configFunc)
	configFunc.EXPECT().FindOne(gomock.Any()).Return(sysConfig, nil)

	configService := NewService(srv, siteDao, categoryDao, configDao)
	_, err := configService.About(ctx)

	assert.NoError(t, err)
}

func TestService_Index(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	siteDao := repository.NewMockIStSiteDao(ctrl)
	siteFunc := repository.NewMockiCustomGenStSiteFunc(ctrl)

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	categoryFunc := repository.NewMockiCustomGenStCategoryFunc(ctrl)

	configDao := repository.NewMockISysConfigDao(ctrl)
	configFunc := repository.NewMockiCustomGenSysConfigFunc(ctrl)

	categories := []*model.StCategory{{ID: 1, ParentID: 0}}
	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)
	categoryDao.EXPECT().WhereByIsUsed(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	categoryFunc.EXPECT().FindAllOrderBySort(gomock.Any(), gomock.Any()).Return(categories, nil)

	sites := []*model.StSite{{}}
	siteDao.EXPECT().WithContext(gomock.Any()).Return(siteFunc)
	siteDao.EXPECT().WhereByIsUsed(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao
	})
	siteFunc.EXPECT().FindAll(gomock.Any()).Return(sites, nil)

	sysConfig := &model.SysConfig{}
	configDao.EXPECT().WithContext(gomock.Any()).Return(configFunc)
	configFunc.EXPECT().FindOne(gomock.Any()).Return(sysConfig, nil)

	configService := NewService(srv, siteDao, categoryDao, configDao)
	_, err := configService.Index(ctx)

	assert.NoError(t, err)
}
