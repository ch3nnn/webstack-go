package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Pacific73/gorm-cache/cache"
	"github.com/Pacific73/gorm-cache/config"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/pkg/log"
	"github.com/ch3nnn/webstack-go/pkg/zapgorm2"
)

type Repository struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewRepository(
	logger *log.Logger,
	db *gorm.DB,
) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

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
	l.Info("db driver", zap.String("driver", dialector.Name()))

	db, err = gorm.Open(dialector, &gorm.Config{
		QueryFields:    true,
		TranslateError: true, // 方言转换错误
		Logger:         zapgorm2.New(l.Logger),
	})
	if err != nil {
		panic(err)
	}

	cachePlugin, _ := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         config.CacheStorageMemory,
		InvalidateWhenUpdate: true,         // when you create/update/delete objects, invalidate cache
		CacheTTL:             3600000 * 24, // 1 day
		CacheSize:            10000,        // max items
	})
	if err := db.Use(cachePlugin); err != nil {
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
	ctx := context.Background()

	err := db.Migrator().DropTable(
		&model.SysUserMenu{},
		&model.SysMenu{},
	)
	if err != nil {
		fmt.Println("migrate drop table error")
		os.Exit(0)
	}

	err = db.AutoMigrate(
		&model.SysConfig{},
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

	_, err = query.SysUser.WithContext(ctx).
		Where(
			query.SysUser.ID.Eq(1),
			query.SysUser.Username.Eq(DefaultUname),
		).
		Attrs(
			field.Attrs(&model.SysUser{
				Password: cryptor.Md5String(DefaultUPassword),
			}),
		).
		FirstOrCreate()
	if err != nil {
		fmt.Println("user migrate error")
		os.Exit(0)
	}

	err = query.SysMenu.WithContext(ctx).
		Create(
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
			&model.SysMenu{
				ID:     4,
				Pid:    0,
				Name:   "系统管理",
				Level:  1,
				Sort:   600,
				IsUsed: true,
			},
			&model.SysMenu{
				ID:     5,
				Pid:    4,
				Name:   "网站配置",
				Link:   "/admin/config",
				Level:  1,
				Sort:   601,
				IsUsed: true,
			},
		)
	if err != nil {
		fmt.Println("menu migrate error")
		os.Exit(0)
	}

	err = query.SysUserMenu.WithContext(ctx).
		Create(
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
			&model.SysUserMenu{
				UserID: 1,
				MenuID: 4,
			},
			&model.SysUserMenu{
				UserID: 1,
				MenuID: 5,
			},
		)
	if err != nil {
		fmt.Println("user menu migrate error")
		os.Exit(0)
	}

	_, err = query.SysConfig.WithContext(ctx).
		Where(
			query.SysConfig.ID.Eq(1),
		).
		Attrs(
			field.Attrs(&model.SysConfig{
				AboutSite:   DefaultAboutSite,
				AboutAuthor: DefaultAuthor,
				SiteTitle:   DefaultSiteTitle,
				SiteKeyword: DefaultSiteKeyword,
				SiteDesc:    DefaultSiteDesc,
			}),
		).
		FirstOrCreate()
	if err != nil {
		fmt.Println("config migrate error")
		os.Exit(0)
	}

	fmt.Println("success initialize")
}
