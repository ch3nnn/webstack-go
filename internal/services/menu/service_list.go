package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/menu"
)

type SearchData struct {
	Pid int32 // 父类ID
}

func (s *service) List(ctx core.Context, searchData *SearchData) (listData []*menu.Menu, err error) {

	qb := menu.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchData.Pid != 0 {
		qb.WherePid(mysql.EqualPredicate, searchData.Pid)
	}

	listData, err = qb.
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
