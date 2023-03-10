package category

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/category"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx core.Context, searchData *SearchData) (listData []*category.Category, err error)
	Tree(ctx core.Context) (nodes []*TreeNode, err error)
	Create(ctx core.Context, siteData *CreateCategoryData) (id int32, err error)
	Modify(ctx core.Context, id int32, menuData *UpdateCategoryData) (err error)
	Delete(ctx core.Context, id int32) error
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *category.Category, err error)
	UpdateUsed(ctx core.Context, id int32, used int32) (err error)
	UpdateSort(ctx core.Context, id int32, sort int32) (err error)
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
