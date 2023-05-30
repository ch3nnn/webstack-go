package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx core.Context) (categories []*model.Category, err error)
	Tree(ctx core.Context) (nodes []*TreeNode, err error)
	Create(ctx core.Context, siteData *CreateCategoryData) (err error)
	Modify(ctx core.Context, id int64, updateCategory *UpdateCategory) (err error)
	Delete(ctx core.Context, id int64) error
	Detail(ctx core.Context, id int64) (category *model.Category, err error)
	UpdateUsed(ctx core.Context, id, used int64) (err error)
	UpdateSort(ctx core.Context, id, sort int64) (err error)
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
