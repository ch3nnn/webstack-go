package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/pkg/log"
	"github.com/ch3nnn/webstack-go/pkg/zapgorm2"
)

type Repository struct {
	logger *log.Logger
	// cache  *cache.Cache
	db *gorm.DB
}

func NewRepository(
	logger *log.Logger,
	// cache *cache.Cache,
	db *gorm.DB,
) *Repository {
	return &Repository{
		logger: logger,
		// cache:  cache,
		db: db,
	}
}

//func NewCache() *cache.Cache {
//	return cache.New(5*time.Minute, 10*time.Minute)
//}

func NewDB(conf *viper.Viper, l *log.Logger) *gorm.DB {
	var (
		db        *gorm.DB
		err       error
		dialector gorm.Dialector
	)

	dsn := conf.GetString("data.db.user.dsn")
	switch conf.GetString("data.db.user.driver") {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":

		dialector = sqlite.Open(dsn)
	default:
		panic("unknown db driver")
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		QueryFields:    true,
		TranslateError: true, // 方言转换错误
		Logger:         zapgorm2.New(l.Logger),
	})
	if err != nil {
		panic(err)
	}

	// Connection Pool config
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// set gorm-gen
	query.SetDefault(db)

	// 迁移和初始化操作
	autoMigrateAndInitialize(db)

	return db.Debug()
}

func autoMigrateAndInitialize(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.SysUser{},
		&model.SysUserMenu{},
		&model.SysMenu{},
		&model.StCategory{},
		&model.StSite{},
	)
	if err != nil {
		fmt.Println("user migrate error")
		os.Exit(0)
	}

	ctx := context.Background()

	umCtn, err := query.SysUserMenu.WithContext(ctx).Count()
	if err != nil {
		os.Exit(0)
	}

	uCtn, err := query.SysUser.WithContext(ctx).Count()
	if err != nil {
		os.Exit(0)
	}

	mCtn, err := query.SysMenu.WithContext(ctx).Count()
	if err != nil {
		os.Exit(0)
	}

	if umCtn == 0 && uCtn == 0 && mCtn == 0 {
		err := query.SysUser.WithContext(ctx).Create(&model.SysUser{
			ID:       1,
			Username: "admin",
			Password: cryptor.Md5String("admin"),
		})
		if err != nil {
			os.Exit(0)
		}

		err = query.SysMenu.WithContext(ctx).Create(
			&model.SysMenu{
				ID:     1,
				Pid:    0,
				Name:   "网站管理",
				Icon:   "users",
				Level:  1,
				Sort:   500,
				IsUsed: true,
			},
			&model.SysMenu{
				ID:     2,
				Pid:    1,
				Name:   "网站分类",
				Link:   "/admin/category",
				Level:  1,
				Sort:   501,
				IsUsed: true,
			},
			&model.SysMenu{
				ID:     3,
				Pid:    1,
				Name:   "网站信息",
				Link:   "/admin/site",
				Level:  1,
				Sort:   502,
				IsUsed: true,
			},
		)
		if err != nil {
			os.Exit(0)
		}

		err = query.SysUserMenu.WithContext(ctx).Create(
			&model.SysUserMenu{
				UserID: 1,
				MenuID: 1,
			},
			&model.SysUserMenu{
				UserID: 1,
				MenuID: 2,
			},
			&model.SysUserMenu{
				UserID: 1,
				MenuID: 3,
			},
		)
		if err != nil {
			os.Exit(0)
		}

		fmt.Println("success initialize")
	}
}
