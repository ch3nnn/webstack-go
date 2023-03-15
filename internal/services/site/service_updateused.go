package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	qb := site.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	data := map[string]any{"IsUsed": used}
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
