/**
 * @Author: chentong
 * @Date: 2025/02/11 14:02
 */

package category

import (
	"flag"
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

func TestService_Create(t *testing.T) {
	ctx := &gin.Context{}

	category := model.StCategory{
		ID:        1,
		ParentID:  1,
		Sort:      1,
		Title:     "webstack-go",
		Icon:      "webstack-go",
		Level:     1,
		IsUsed:    true,
		CreatedAt: nil,
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	categoryFunc := repository.NewMockiCustomGenStCategoryFunc(ctrl)

	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)
	categoryFunc.EXPECT().Create(gomock.Any()).Return(&category, nil)

	categoryService := NewService(srv, categoryDao)
	_, err := categoryService.Create(ctx, &v1.CategoryCreateReq{
		ParentID: 1,
		Level:    1,
		Name:     "webstack-go",
		Icon:     "webstack-go",
		IsUsed:   true,
		SortID:   1,
	})

	assert.NoError(t, err)
}

func TestService_Delete(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	categoryFunc := repository.NewMockiCustomGenStCategoryFunc(ctrl)

	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)
	categoryDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao { return dao })
	categoryFunc.EXPECT().Delete(gomock.Any()).Return(nil)

	categoryService := NewService(srv, categoryDao)
	_, err := categoryService.Delete(ctx, &v1.CategoryDeleteReq{ID: 1})

	assert.NoError(t, err)
}

func TestService_List(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	categoryFunc := repository.NewMockiCustomGenStCategoryFunc(ctrl)

	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)

	categories := []*model.StCategory{{}}
	categoryFunc.EXPECT().FindAllOrderBySort(gomock.Any(), gomock.Any()).Return(categories, nil)

	categoryService := NewService(srv, categoryDao)
	_, err := categoryService.List(ctx, &v1.CategoryListReq{})

	assert.NoError(t, err)
}

func TestService_Detail(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	categoryFunc := repository.NewMockiCustomGenStCategoryFunc(ctrl)

	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)
	categoryDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao { return dao })

	category := &model.StCategory{
		ID:       1,
		ParentID: 1,
		Sort:     1,
		Title:    "webstack-go",
		Icon:     "webstack-go",
		Level:    1,
		IsUsed:   false,
	}
	categoryFunc.EXPECT().FindOne(gomock.Any()).Return(category, nil)

	categoryService := NewService(srv, categoryDao)
	_, err := categoryService.Detail(ctx, &v1.CategoryDetailReq{ID: 1})

	assert.NoError(t, err)
}

func TestService_Update(t *testing.T) {
	ctx := &gin.Context{}

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	categoryFunc := repository.NewMockiCustomGenStCategoryFunc(ctrl)

	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)
	categoryDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao { return dao })
	categoryFunc.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	category := &model.StCategory{
		ID:       1,
		ParentID: 1,
		Sort:     1,
		Title:    "webstack-go",
		Icon:     "webstack-go",
		Level:    1,
		IsUsed:   false,
	}
	categoryDao.EXPECT().WithContext(gomock.Any()).Return(categoryFunc)
	categoryDao.EXPECT().WhereByID(gomock.Any()).Return(func(dao gen.Dao) gen.Dao { return dao })
	categoryFunc.EXPECT().FindOne(gomock.Any()).Return(category, nil)

	categoryService := NewService(srv, categoryDao)
	_, err := categoryService.Update(ctx, &v1.CategoryUpdateReq{
		ID:     1,
		Pid:    pointer.Of(1),
		Name:   pointer.Of("name"),
		Icon:   pointer.Of("icon"),
		IsUsed: pointer.Of(true),
		SortID: pointer.Of(1),
	})

	assert.NoError(t, err)
}
