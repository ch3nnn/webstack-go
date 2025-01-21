package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

// func NewCache() *cache.Cache {
//	return cache.New(5*time.Minute, 10*time.Minute)
// }

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
			os.Exit(0)
		}

		err = query.SysConfig.WithContext(ctx).Create(
			&model.SysConfig{
				ID:        1,
				AboutSite: "> ❤️ 基于 Golang 开源的网址导航网站项目，具备完整的前后台，您可以拿来制作自己平日收藏的网址导航。\n\n\n> 如果你也是开发者，如果你也正好喜欢折腾，那希望这个网站能给你带来一些作用。",
				AboutAuthor: `
<div class="col-sm-4">
    <div class="xe-widget xe-conversations box2 label-info" onclick="window.open('https://blog.ch3nnn.cn/about/', '_blank')" data-toggle="tooltip" data-placement="bottom" title="" data-original-title="https://blog.ch3nnn.cn/about/">
        <div class="xe-comment-entry">
            <a class="xe-user-img">
                <img src="https://s2.loli.net/2023/02/20/H1k52mlXNYKWDrU.png" class="img-circle" width="40">
            </a>
            <div class="xe-comment">
                <a href="#" class="xe-user-name overflowClip_1">
                    <strong>Developer. Ch3nnn.</strong>
                </a>
                <p class="overflowClip_2"> 折腾不息 · 乐此不疲.</p>
            </div>
        </div>
    </div>
</div>

<div class="col-md-8">
    <div class="row">
        <div class="col-sm-12">
            <br>
            <blockquote>
                <p>
                    这是一个公益项目，而且是<a href="https://github.com/ch3nnn/webstack-go"> 开源 </a>的。你也可以拿来制作自己的网址导航。如果你有更好的想法，可以通过个人网站<a href="https://ch3nnn.cn/about/">ch3nnn.cn</a>中的联系方式找到我，欢迎与我交流分享。
                </p>
            </blockquote>
        </div>
    </div>
    <br>
</div>
`,
				IsAbout:     false,
				SiteTitle:   "WebStack-Go - 网址导航",
				SiteKeyword: "网址导航",
				SiteDesc:    "WebStack-Go - 基于 Golang 开源的网址导航网站",
			})
		if err != nil {
			os.Exit(0)
		}

		fmt.Println("success initialize")
	}
}
