package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"gorm.io/gorm"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {
	// 先查询 id 是否存在
	_, err = site.NewQueryBuilder().
		WhereId(mysql.EqualPredicate, id).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	qb := site.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	if err = qb.Delete(s.db.GetDbW().WithContext(ctx.RequestContext())); err != nil {
		return err
	}

	return
}
