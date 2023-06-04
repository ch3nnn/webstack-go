package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchMyMenuData struct {
	AdminId int64 `json:"admin_id"` // 管理员ID
}

type ListMyMenuData struct {
	Id   int64  `json:"id"`   // ID
	Pid  int64  `json:"pid"`  // 父类ID
	Name string `json:"name"` // 菜单名称
	Link string `json:"link"` // 链接地址
	Icon string `json:"icon"` // 图标
}

func (s *service) MyMenu(ctx core.Context, searchData *SearchMyMenuData) (menuData []ListMyMenuData, err error) {
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

	for _, menu := range menus {
		for _, adminMenu := range adminMenus {
			if menu.ID == adminMenu.MenuID {
				menuData = append(menuData, ListMyMenuData{
					Id:   menu.ID,
					Pid:  menu.Pid,
					Name: menu.Name,
					Link: menu.Link,
					Icon: menu.Icon,
				})
			}
		}
	}

	return
}
