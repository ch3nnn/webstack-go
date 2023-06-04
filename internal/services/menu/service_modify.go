package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type UpdateMenuData struct {
	Name string // 菜单名称
	Link string // 链接地址
	Icon string // 图标
}

func (s *service) Modify(ctx core.Context, id int64, menuData *UpdateMenuData) (err error) {

	if _, err = query.Menu.WithContext(ctx.RequestContext()).
		Where(query.Menu.ID.Eq(id)).
		UpdateColumnSimple(
			query.Menu.Name.Value(menuData.Name),
			query.Menu.Link.Value(menuData.Link),
			query.Menu.Icon.Value(menuData.Icon),
			query.Menu.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	return
}
