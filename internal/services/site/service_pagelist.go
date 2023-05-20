package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
)

type SearchData struct {
	Page              int    `json:"page"`               // 第几页
	PageSize          int    `json:"page_size"`          // 每页显示条数
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessSecret    string `json:"business_secret"`    // 调用方secret
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
	Search            string `json:"search"`             // 搜索关键字
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*site.Site, err error) {

	page := searchData.Page
	if page == 0 {
		page = 1
	}

	pageSize := searchData.PageSize
	if pageSize == 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := site.NewQueryBuilder()
	if searchData.Search != "" {
		qb = qb.WhereTitleLike(searchData.Search)
	}
	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
