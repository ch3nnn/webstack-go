/**
 * @Author: chentong
 * @Date: 2025/02/08 15:23
 */

package user

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
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
	"github.com/ch3nnn/webstack-go/internal/middleware"
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

func TestUserService_Info(t *testing.T) {
	user := &model.SysUser{
		ID:        1,
		Username:  "admin",
		Password:  "admin",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}
	menus := repository.DefaultSysMenuAdmin
	userMenus := repository.DefaultSysUserMenuAdmin

	ctx := &gin.Context{}
	ctx.Set(middleware.UserID, user.ID)

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	userDao := repository.NewMockISysUserDao(ctrl)
	siteDao := repository.NewMockIStSiteDao(ctrl)
	menuDao := repository.NewMockISysMenuDao(ctrl)
	userMenuDao := repository.NewMockISysUserMenuDao(ctrl)

	// user repo
	userFunc := repository.NewMockiCustomGenSysUserFunc(ctrl)
	userDao.EXPECT().WithContext(gomock.Any()).Return(userFunc)
	userDao.EXPECT().WhereByID(user.ID).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.ID.Eq(user.ID))
	})
	userFunc.EXPECT().FindOne(gomock.Any()).Return(user, nil)
	// menu repo
	menuFunc := repository.NewMockiCustomGenSysMenuFunc(ctrl)
	menuDao.EXPECT().WithContext(gomock.Any()).Return(menuFunc)
	menuFunc.EXPECT().FindAll().Return(menus, nil)
	// usermenu repo
	userMenuFunc := repository.NewMockiCustomGenSysUserMenuFunc(ctrl)
	userMenuDao.EXPECT().WithContext(gomock.Any()).Return(userMenuFunc)
	userMenuDao.EXPECT().WhereByUserID(user.ID).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUserMenu.UserID.Eq(user.ID))
	})
	userMenuFunc.EXPECT().FindAll(gomock.Any()).Return(userMenus, nil)

	userService := NewService(srv, userDao, siteDao, categoryDao, menuDao, userMenuDao)
	infoResp, err := userService.Info(ctx, nil)

	// 断言
	assert.NoError(t, err)
	assert.Equal(t, user.Username, infoResp.Username)
}

func TestUserService_Login(t *testing.T) {
	ctx := context.Background()

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	userDao := repository.NewMockISysUserDao(ctrl)
	siteDao := repository.NewMockIStSiteDao(ctrl)
	menuDao := repository.NewMockISysMenuDao(ctrl)
	userMenuDao := repository.NewMockISysUserMenuDao(ctrl)

	user := &model.SysUser{
		ID:        1,
		Username:  "admin",
		Password:  "admin",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}
	req := &v1.LoginReq{Username: user.Username, Password: user.Password}

	userFunc := repository.NewMockiCustomGenSysUserFunc(ctrl)

	userDao.EXPECT().WithContext(gomock.Any()).Return(userFunc)
	userDao.EXPECT().WhereByPassword(gomock.Any()).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.Password.Eq(req.Password))
	})
	userDao.EXPECT().WhereByUsername(user.Username).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.Password.Eq(req.Username))
	})

	userFunc.EXPECT().FindOne(gomock.Any()).Return(user, nil)

	userService := NewService(srv, userDao, siteDao, categoryDao, menuDao, userMenuDao)
	loginResp, err := userService.Login(ctx, req)
	assert.NoError(t, err)

	token, err := j.GenToken(user.ID, time.Now().Add(time.Hour*24))
	assert.NoError(t, err)

	assert.Equal(t, token, loginResp.Token)
}

func TestService_UpdatePassword(t *testing.T) {
	user := &model.SysUser{
		ID:        1,
		Username:  "admin",
		Password:  "admin",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}

	req := &v1.UpdatePasswordReq{
		OldPassword: "admin",
		NewPassword: "admin",
	}

	ctx := &gin.Context{}
	ctx.Set(middleware.UserID, user.ID)

	repo, _ := setupRepository(t)
	srv := s.NewService(logger, j, repo)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	categoryDao := repository.NewMockIStCategoryDao(ctrl)
	userDao := repository.NewMockISysUserDao(ctrl)
	siteDao := repository.NewMockIStSiteDao(ctrl)
	menuDao := repository.NewMockISysMenuDao(ctrl)
	userMenuDao := repository.NewMockISysUserMenuDao(ctrl)

	// find one
	userFunc := repository.NewMockiCustomGenSysUserFunc(ctrl)
	userDao.EXPECT().WithContext(gomock.Any()).Return(userFunc)
	userDao.EXPECT().WhereByID(user.ID).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.ID.Eq(user.ID))
	})
	userFunc.EXPECT().FindOne(gomock.Any()).Return(user, nil)
	// update
	userDao.EXPECT().WithContext(gomock.Any()).Return(userFunc)
	userDao.EXPECT().WhereByID(user.ID).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.ID.Eq(user.ID))
	})
	userFunc.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(1), nil)

	// 调用被测试的方法
	userService := NewService(srv, userDao, siteDao, categoryDao, menuDao, userMenuDao)
	resp, err := userService.UpdatePassword(ctx, req)

	// 断言
	assert.NoError(t, err)
	assert.Empty(t, resp)

	// error
	userDao.EXPECT().WithContext(gomock.Any()).Return(userFunc)
	userDao.EXPECT().WhereByID(user.ID).Return(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.ID.Eq(user.ID))
	})
	userFunc.EXPECT().FindOne(gomock.Any()).Return(nil, gorm.ErrRecordNotFound)

	resp, err = userService.UpdatePassword(ctx, req)
	// 断言
	assert.Error(t, err)
	assert.Empty(t, resp)
}

func TestSysUserDao_FindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userFunc := repository.NewMockiCustomGenSysUserFunc(ctrl)

	user := &model.SysUser{
		ID:        1,
		Username:  "admin",
		Password:  "admin",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}

	userFunc.EXPECT().FindOne(gomock.Any()).Return(user, nil)

	result, err := userFunc.FindOne(func(dao gen.Dao) gen.Dao {
		return dao.Where(query.SysUser.ID.Eq(1))
	})

	assert.NoError(t, err)
	assert.Equal(t, user, result)
}
