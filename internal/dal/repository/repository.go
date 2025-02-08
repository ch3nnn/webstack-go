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
	"github.com/ch3nnn/webstack-go/pkg/gormx"
	"github.com/ch3nnn/webstack-go/pkg/log"
	"github.com/ch3nnn/webstack-go/pkg/zapgorm2"
)

type Repository struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
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
	case gormx.MYSQL:
		dialector = mysql.Open(dsn)
	case gormx.POSTGRES:
		dialector = postgres.Open(dsn)
	case gormx.SQLITE:
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
		InvalidateWhenUpdate: true,  // when you create/update/delete objects, invalidate cache
		CacheTTL:             30000, // 30 s
		CacheSize:            10000, // max items
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
		).
		Attrs(
			field.Attrs(&model.SysUser{
				Username: DefaultUname,
				Password: cryptor.Md5String(DefaultUPassword),
			}),
		).
		FirstOrCreate()
	if err != nil {
		fmt.Println("user migrate error")
		os.Exit(0)
	}

	if err = query.SysMenu.WithContext(ctx).Create(DefaultSysMenuAdmin...); err != nil {
		fmt.Println("menu migrate error")
		os.Exit(0)
	}

	if err = query.SysUserMenu.WithContext(ctx).Create(DefaultSysUserMenuAdmin...); err != nil {
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
				SiteFavicon: DefaultFaviconBase64,
				SiteLogo:    DefaultLogoBase64,
			}),
		).
		FirstOrCreate()
	if err != nil {
		fmt.Println("config migrate error")
		os.Exit(0)
	}

	fmt.Println("success initialize")
}
