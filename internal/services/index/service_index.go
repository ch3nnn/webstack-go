package index

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
)

func (s *service) Index(ctx core.Context) (listData []*site.Site, err error) {

	qb := site.NewQueryBuilder()

	listData, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
