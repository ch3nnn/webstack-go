package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Create(ctx core.Context, adminData *CreateAdminData) (id int64, err error)
	PageList(ctx core.Context, searchData *SearchData) (admins []*model.Admin, err error)
	PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error)
	UpdateUsed(ctx core.Context, id, used int64) (err error)
	Delete(ctx core.Context, id int64) (err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *model.Admin, err error)
	ResetPassword(ctx core.Context, id int64) (err error)
	ModifyPassword(ctx core.Context, id int64, newPassword string) (err error)
	ModifyPersonalInfo(ctx core.Context, id int64, modifyData *ModifyData) (err error)

	CreateMenu(ctx core.Context, menuData *CreateMenuData) (err error)
	ListMenu(ctx core.Context, searchData *SearchListMenuData) (menuData []ListMenuData, err error)
	MyMenu(ctx core.Context, searchData *SearchMyMenuData) (menuData []ListMyMenuData, err error)
	MyAction(ctx core.Context, searchData *SearchMyActionData) (actionData []MyActionData, err error)
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
