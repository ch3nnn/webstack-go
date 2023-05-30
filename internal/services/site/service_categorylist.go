package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"sort"
	"strconv"
	"strings"
)

func stringSliceToInt64Slice(stringSlice []string) (int64Slice []int64, err error) {
	for _, str := range stringSlice {
		if i, err := strconv.ParseInt(str, 10, 64); err != nil {
			return nil, err
		} else {
			int64Slice = append(int64Slice, i)
		}
	}
	return int64Slice, nil
}

func (s *service) CategoryList(ctx core.Context) (categories []*model.Category, err error) {
	// 查询父 id
	result, err := query.Category.WithContext(ctx.RequestContext()).GetParentIdsByGroupParentId()
	if err != nil {
		return nil, err
	}
	stringSlice := strings.Split(result["parent_ids"].(string), ",")
	parentIds, err := stringSliceToInt64Slice(stringSlice)
	if err != nil {
		return nil, err
	}
	// 一级分类
	categories01, err := query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ParentID.In(parentIds...)).
		Not(query.Category.ParentID.Eq(0)).
		Find()
	if err != nil {
		return nil, err
	}
	// 二级分类
	categories02, err := query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ParentID.Eq(0)).
		Not(query.Category.ID.In(parentIds...)).
		Find()
	if err != nil {
		return nil, err
	}

	categories = append(categories01, categories02...)
	// 按分类升序
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Sort < categories[j].Sort
	})
	return
}
