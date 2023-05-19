package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"

	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
)

func (s *service) Delete(ctx core.Context, id int32) (err error) {

	qb1 := category.NewQueryBuilder().WhereId(mysql.EqualPredicate, id)
	err = qb1.Delete(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	// 删除一级目录 id 需要删除二级分类
	qb2 := category.NewQueryBuilder().WhereParentId(mysql.EqualPredicate, id)
	err = qb2.Delete(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}

	return
}
