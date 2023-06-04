package authorized

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int64, err error)
	List(ctx core.Context, searchData *SearchData) (authorizedList []*model.Authorized, err error)
	PageList(ctx core.Context, searchData *SearchData) (authorizedList []*model.Authorized, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id, used int64) (err error)
	Delete(ctx core.Context, id int64) (err error)
	Detail(ctx core.Context, id int64) (authorizedInfo *model.Authorized, err error)
	DetailByKey(ctx core.Context, key string) (data *CacheAuthorizedData, err error)

	CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int64, err error)
	ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (authorizedAPIS []*model.AuthorizedAPI, err error)
	DeleteAPI(ctx core.Context, id int64) (err error)
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
