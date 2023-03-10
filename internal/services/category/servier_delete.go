package category

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"

	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/category"
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
