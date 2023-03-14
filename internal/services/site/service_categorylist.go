package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
	"sort"
)

func (s *service) CategoryList(ctx core.Context) (listData []*category.Category, err error) {
	qb := category.NewQueryBuilder()
	parentIds := qb.GroupByParentId(s.db.GetDbR().WithContext(ctx.RequestContext()))
	// 一级分类
	qb1 := category.NewQueryBuilder()
	categories01, err := qb1.
		WhereParentIdIn(parentIds).
		WhereParentId(mysql.NotEqualPredicate, 0).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	// 二级分类
	qb2 := category.NewQueryBuilder()
	categories02, err := qb2.
		WhereIdNotIn(parentIds).
		WhereParentId(mysql.EqualPredicate, 0).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	listData = append(categories01, categories02...)
	// 按分类升序
	sort.Slice(listData, func(i, j int) bool {
		return listData[i].Sort < listData[j].Sort
	})
	return
}
