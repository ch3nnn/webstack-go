package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) PageListCount(ctx core.Context, searchData *SearchData) (total int64, err error) {

	iSiteDo := query.Site.WithContext(ctx.RequestContext())
	if searchData.Search != "" {
		iSiteDo = iSiteDo.Where(query.Site.Title.Like("%" + searchData.Search + "%"))
	}
	if total, err = iSiteDo.Count(); err != nil {
		return 0, err
	} else {
		return
	}
}
