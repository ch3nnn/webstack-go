package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) UpdateUsed(ctx core.Context, id, used int64) (err error) {
	if _, err = query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.ID.Eq(id)).
		Update(query.Site.IsUsed, used); err != nil {
		return err
	}

	return
}
