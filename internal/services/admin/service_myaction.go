package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchMyActionData struct {
	AdminId int64 `json:"admin_id"` // 管理员ID
}

type MyActionData struct {
	Id     int64  // 主键
	MenuId int64  // 菜单栏ID
	Method string // 请求方式
	Api    string // 请求地址
}

func (s *service) MyAction(ctx core.Context, searchData *SearchMyActionData) (actionData []MyActionData, err error) {
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

	var menuIds []int64
	for _, v := range adminMenus {
		menuIds = append(menuIds, v.MenuID)
	}

	menuActions, err := query.MenuAction.WithContext(ctx.RequestContext()).
		Where(query.MenuAction.IsDeleted.Eq(-1)).
		Where(query.MenuAction.ID.In(menuIds...)).
		Find()
	if err != nil {
		return nil, err
	}

	if len(menuActions) <= 0 {
		return
	}

	actionData = make([]MyActionData, len(menuActions))
	for k, v := range menuActions {
		actionData[k] = MyActionData{
			Id:     v.ID,
			MenuId: v.MenuID,
			Method: v.Method,
			Api:    v.API,
		}
	}

	return
}
