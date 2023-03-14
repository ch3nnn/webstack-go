package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
)

type CategorySite struct {
	Category string
	SiteList []site.Site
}

func (s *service) CategorySite(ctx core.Context) (categorySites []*CategorySite, err error) {
	// 获取分类
	categoryQueryBuilder := category.NewQueryBuilder()
	categories, err := categoryQueryBuilder.
		OrderBySort(true).
		OrderById(true).
		WhereIsUsed(mysql.EqualPredicate, 1).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	// 获取网站信息
	siteQueryBuilder := site.NewQueryBuilder()
	sites, err := siteQueryBuilder.
		WhereIsUsed(mysql.EqualPredicate, 1).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	// 拼接数据
	for _, ctg := range categories {
		var siteList []site.Site
		for _, st := range sites {
			if ctg.Id == st.CategoryId {
				siteList = append(siteList, *st)
			}
		}
		categorySites = append(categorySites, &CategorySite{
			Category: ctg.Title,
			SiteList: siteList,
		})
	}

	return
}
