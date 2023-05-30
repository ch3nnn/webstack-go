package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type CreateCategoryData struct {
	Pid   int64  // 父类ID
	Name  string // 分类名称
	Icon  string // 图标库 https://lineicons.com/icons/
	Level int64
}

func (s *service) Create(ctx core.Context, siteData *CreateCategoryData) (err error) {
	if err = query.Category.WithContext(ctx.RequestContext()).
		Create(&model.Category{
			ParentID: siteData.Pid,
			Title:    siteData.Name,
			Icon:     siteData.Icon,
			Level:    siteData.Level,
		}); err != nil {
		return
	}
	return
}
