package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchListMenuData struct {
	AdminId int64 `json:"admin_id"` // 管理员ID
}

type ListMenuData struct {
	Id     int64  `json:"id"`      // ID
	Pid    int64  `json:"pid"`     // 父类ID
	Name   string `json:"name"`    // 菜单名称
	IsHave int64  `json:"is_have"` // 是否已拥有权限
}

func (s *service) ListMenu(ctx core.Context, searchData *SearchListMenuData) (menuData []ListMenuData, err error) {
	menus, err := query.Menu.WithContext(ctx.RequestContext()).
		Where(query.Menu.IsDeleted.Eq(-1)).
		Order(query.Menu.Sort).
		Find()
	if err != nil {
		return nil, err
	}

	if len(menus) <= 0 {
		return
	}

	menuData = make([]ListMenuData, len(menus))
	for k, v := range menus {
		menuData[k] = ListMenuData{
			Id:     v.ID,
			Pid:    v.Pid,
			Name:   v.Name,
			IsHave: 0,
		}
	}

	iAdminMenuDo := query.AdminMenu.WithContext(ctx.RequestContext())
	if searchData.AdminId != 0 {
		iAdminMenuDo = iAdminMenuDo.Where(query.AdminMenu.AdminID.Eq(searchData.AdminId))
	}
	adminMenus, err := iAdminMenuDo.Order(query.AdminMenu.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}

	if len(adminMenus) <= 0 {
		return
	}

	for k, v := range menuData {
		for _, haveV := range adminMenus {
			if haveV.MenuID == v.Id {
				menuData[k].IsHave = 1
			}
		}
	}

	return
}
