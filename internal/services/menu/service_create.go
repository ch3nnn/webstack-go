package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"time"
)

type CreateMenuData struct {
	Pid   int64  // 父类ID
	Name  string // 菜单名称
	Link  string // 链接地址
	Icon  string // 图标
	Level int64  // 菜单类型 1:一级菜单 2:二级菜单
}

func (s *service) Create(ctx core.Context, menuData *CreateMenuData) (id int64, err error) {
	if err = query.Menu.WithContext(ctx.RequestContext()).Create(&model.Menu{
		Pid:         menuData.Pid,
		Name:        menuData.Name,
		Link:        menuData.Link,
		Icon:        menuData.Icon,
		Level:       menuData.Level,
		IsUsed:      1,
		IsDeleted:   -1,
		CreatedAt:   time.Time{},
		CreatedUser: ctx.SessionUserInfo().UserName,
	}); err != nil {
		return 0, err
	}
	return
}
