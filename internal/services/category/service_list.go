package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
)

type SearchData struct {
	Pid int32 // 父类ID
}

func (s *service) List(ctx core.Context, searchData *SearchData) (listData []*category.Category, err error) {

	qb := category.NewQueryBuilder()
	listData, err = qb.
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
