package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"

	"gorm.io/gorm"
)

func (s *service) DeleteAction(ctx core.Context, id int64) (err error) {
	// 先查询 id 是否存在
	_, err = query.MenuAction.WithContext(ctx.RequestContext()).
		Where(query.MenuAction.IsDeleted.Eq(-1)).
		Where(query.MenuAction.ID.Eq(id)).
		First()
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	if _, err = query.MenuAction.WithContext(ctx.RequestContext()).
		Where(query.MenuAction.ID.Eq(id)).
		UpdateColumnSimple(
			query.MenuAction.IsDeleted.Value(1),
			query.MenuAction.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	return
}
