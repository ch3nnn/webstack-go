package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) Delete(ctx core.Context, id int64) (err error) {
	if _, err = query.Menu.WithContext(ctx.RequestContext()).
		Where(query.Menu.ID.Eq(id)).
		UpdateSimple(
			query.Menu.IsDeleted.Value(1),
			query.Menu.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	return
}
