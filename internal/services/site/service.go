package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx core.Context) (listData []*site.Site, err error)
	CategoryList(ctx core.Context) (listData []*category.Category, err error)
	PageList(ctx core.Context, searchData *SearchData) (listData []*site.Site, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	Delete(ctx core.Context, id int32) (err error)
	Create(ctx core.Context, siteData *CreateSiteData) (id int32, err error)
	CategorySite(ctx core.Context) (categorySites []*CategorySite, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
