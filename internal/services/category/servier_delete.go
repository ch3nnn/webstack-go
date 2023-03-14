package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"

	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	qb := category.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Delete(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	return
}
