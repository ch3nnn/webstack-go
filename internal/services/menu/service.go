package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, menuData *CreateMenuData) (id int64, err error)
	Modify(ctx core.Context, id int64, menuData *UpdateMenuData) (err error)
	List(ctx core.Context, searchData *SearchData) (menus []*model.Menu, err error)
	UpdateUsed(ctx core.Context, id, sort int64) (err error)
	UpdateSort(ctx core.Context, id, sort int64) (err error)
	Delete(ctx core.Context, id int64) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (menu *model.Menu, err error)

	CreateAction(ctx core.Context, menuActionData *CreateMenuActionData) (id int32, err error)
	ListAction(ctx core.Context, searchData *SearchListActionData) (menuActions []*model.MenuAction, err error)
	DeleteAction(ctx core.Context, id int64) (err error)
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
