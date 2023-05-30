package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type CategorySite struct {
	Category string
	SiteList []model.Site
}

func (s *service) CategorySite(ctx core.Context) (categorySites []*CategorySite, err error) {
	// 获取分类
	categories, err := query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.IsUsed.Eq(1)).
		Order(query.Category.Sort).
		Order(query.Category.ID).
		Find()

	if err != nil {
		return nil, err
	}
	// 获取网站信息
	sites, err := query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.IsUsed.Eq(1)).
		Find()
	if err != nil {
		return nil, err
	}
	// 拼接数据
	for _, cat := range categories {
		var siteList []model.Site
		for _, st := range sites {
			if cat.ID == st.CategoryID {
				siteList = append(siteList, *st)
			}
		}
		categorySites = append(categorySites, &CategorySite{
			Category: cat.Title,
			SiteList: siteList,
		})
	}

	return
}
