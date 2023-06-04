package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Page     int64  // 第几页
	PageSize int64  // 每页显示条数
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (admins []*model.Admin, err error) {

	iAdminDo := query.Admin.WithContext(ctx.RequestContext())
	iAdminDo = iAdminDo.Where(query.Admin.IsDeleted.Eq(-1))
	if searchData.Username != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Username.Eq(searchData.Username))
	}
	if searchData.Nickname != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Nickname.Eq(searchData.Nickname))
	}
	if searchData.Mobile != "" {
		iAdminDo = iAdminDo.Where(query.Admin.Mobile.Eq(searchData.Mobile))
	}

	admins, _, err = iAdminDo.Order(query.Admin.ID.Desc()).
		FindByPage(int((searchData.Page-1)*searchData.PageSize), int(searchData.PageSize))
	if err != nil {
		return nil, err
	}

	return
}
