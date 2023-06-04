package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"github.com/spf13/cast"
	"strings"
)

type CreateMenuData struct {
	AdminId int64  `form:"admin_id"` // AdminID
	Actions string `form:"actions"`  // 功能权限ID,多个用,分割
}

func (s *service) CreateMenu(ctx core.Context, menuData *CreateMenuData) (err error) {
	if _, err = query.AdminMenu.WithContext(ctx.RequestContext()).
		Where(query.AdminMenu.AdminID.Eq(menuData.AdminId)).
		Delete(); err != nil {
		return err
	}

	menus := make([]*model.AdminMenu, 0, 10)
	for _, v := range strings.Split(menuData.Actions, ",") {
		menus = append(menus, &model.AdminMenu{
			AdminID:     menuData.AdminId,
			MenuID:      cast.ToInt64(v),
			CreatedUser: ctx.SessionUserInfo().UserName,
		})
	}
	if err = query.AdminMenu.CreateInBatches(menus, len(menus)); err != nil {
		return err
	}

	return
}
