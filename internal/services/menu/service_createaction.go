package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"time"
)

type CreateMenuActionData struct {
	MenuId int64  `json:"menu_id"` // 菜单栏ID
	Method string `json:"method"`  // 请求方法
	API    string `json:"api"`     // 请求地址
}

func (s *service) CreateAction(ctx core.Context, menuActionData *CreateMenuActionData) (id int32, err error) {
	if err = query.MenuAction.WithContext(ctx.RequestContext()).
		Create(&model.MenuAction{
			MenuID:      menuActionData.MenuId,
			Method:      menuActionData.Method,
			API:         menuActionData.API,
			IsDeleted:   -1,
			CreatedAt:   time.Time{},
			CreatedUser: ctx.SessionUserInfo().UserName,
		}); err != nil {
		return 0, err
	}

	return
}
