package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) UpdateUsed(ctx core.Context, id, used int64) (err error) {
	if _, err = query.Category.WithContext(ctx.RequestContext()).
		Where(query.Category.ID.Eq(id)).
		Update(query.Category.IsUsed, used); err != nil {
		return err
	}

	return
}
