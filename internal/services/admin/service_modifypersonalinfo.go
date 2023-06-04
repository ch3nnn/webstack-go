package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type ModifyData struct {
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) ModifyPersonalInfo(ctx core.Context, id int64, modifyData *ModifyData) (err error) {

	if _, err = query.Admin.WithContext(ctx.RequestContext()).
		Where(query.Admin.ID.Eq(id)).
		UpdateColumnSimple(
			query.Admin.Nickname.Value(modifyData.Nickname),
			query.Admin.Mobile.Value(modifyData.Mobile),
			query.Admin.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	return
}
