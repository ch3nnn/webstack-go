package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchListActionData struct {
	MenuId int64 `json:"menu_id"` // 菜单栏ID
}

func (s *service) ListAction(ctx core.Context, searchData *SearchListActionData) (menuActions []*model.MenuAction, err error) {
	iMenuActionDo := query.MenuAction.WithContext(ctx.RequestContext())
	iMenuActionDo = iMenuActionDo.Where(query.MenuAction.IsDeleted.Eq(-1))
	if searchData.MenuId != 0 {
		iMenuActionDo = iMenuActionDo.Where(query.MenuAction.MenuID.Eq(searchData.MenuId))
	}
	menuActions, err = iMenuActionDo.Order(query.MenuAction.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}

	return
}
