package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchOneData struct {
	Id       int64  // 用户ID
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
	IsUsed   int64  // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *model.Admin, err error) {
	iAdminDo := query.Admin.WithContext(ctx.RequestContext())
	iAdminDo = iAdminDo.Where(query.Admin.IsDeleted.Eq(-1))

	if searchOneData.Id != 0 {
		iAdminDo = iAdminDo.Where(query.Admin.ID.Eq(searchOneData.Id))
	}

	if searchOneData.Username != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Username.Eq(searchOneData.Username))
	}

	if searchOneData.Nickname != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Nickname.Eq(searchOneData.Nickname))
	}

	if searchOneData.Mobile != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Mobile.Eq(searchOneData.Mobile))
	}

	if searchOneData.Password != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Password.Eq(searchOneData.Password))
	}

	if searchOneData.IsUsed != 0 {
		iAdminDo = iAdminDo.Where(query.Admin.IsUsed.Eq(searchOneData.IsUsed))
	}

	info, err = iAdminDo.First()
	if err != nil {
		return nil, err
	}

	return
}
