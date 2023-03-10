package site

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/site"
	"time"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	qb := site.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	data := map[string]any{
		"IsUsed":     used,
		"UpdateTime": time.Now(),
	}
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
