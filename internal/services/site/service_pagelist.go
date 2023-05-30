package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Page              int64  `json:"page"`               // 第几页
	PageSize          int64  `json:"page_size"`          // 每页显示条数
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessSecret    string `json:"business_secret"`    // 调用方secret
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
	Search            string `json:"search"`             // 搜索关键字
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (sites []*model.Site, err error) {

	iSiteDo := query.Site.WithContext(ctx.RequestContext())
	if searchData.Search != "" {
		iSiteDo = iSiteDo.Where(query.Site.Title.Like("%" + searchData.Search + "%"))
	}
	sites, _, err = iSiteDo.Preload(query.Site.Category).
		Order(query.Site.ID.Desc()).
		FindByPage(int((searchData.Page-1)*searchData.PageSize), int(searchData.PageSize))
	if err != nil {
		return nil, err
	}

	return
}
