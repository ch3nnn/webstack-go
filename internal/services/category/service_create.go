package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
)

type CreateCategoryData struct {
	Pid   int32  // 父类ID
	Name  string // 分类名称
	Icon  string // 图标库 https://lineicons.com/icons/
	Level int32
}

func (s *service) Create(ctx core.Context, siteData *CreateCategoryData) (id int32, err error) {
	model := category.NewModel()
	model.ParentId = siteData.Pid
	model.Title = siteData.Name
	model.Icon = siteData.Icon
	model.Level = siteData.Level

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
