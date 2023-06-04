package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) UpdateUsed(ctx core.Context, id, used int64) (err error) {
	if _, err = query.Menu.WithContext(ctx.RequestContext()).
		Where(query.Menu.ID.Eq(id)).
		UpdateColumnSimple(
			query.Menu.IsUsed.Value(used),
			query.Menu.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	return
}
