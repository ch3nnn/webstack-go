package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
)

type SearchOneData struct {
	Id     int32 // 用户ID
	IsUsed int32 // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *category.Category, err error) {

	qb := category.NewQueryBuilder()

	if searchOneData.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, searchOneData.Id)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
