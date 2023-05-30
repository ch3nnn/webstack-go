package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"gorm.io/gorm"
)

func (s *service) Delete(ctx core.Context, id int64) (err error) {
	// 先查询 id 是否存在
	if _, err = query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.ID.Eq(id)).
		First(); err == gorm.ErrRecordNotFound {
		return nil
	}
	// 根据 id 删除
	if _, err = query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.ID.Eq(id)).
		Delete(); err != nil {
		return
	}

	return
}
