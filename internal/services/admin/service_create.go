package admin

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/password"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type CreateAdminData struct {
	Username string // 用户名
	Nickname string // 昵称
	Mobile   string // 手机号
	Password string // 密码
}

func (s *service) Create(ctx core.Context, adminData *CreateAdminData) (id int64, err error) {
	err = query.Admin.WithContext(ctx.RequestContext()).Create(&model.Admin{
		Username:  adminData.Username,
		Password:  password.GeneratePassword(adminData.Password),
		Nickname:  adminData.Nickname,
		Mobile:    adminData.Mobile,
		IsUsed:    1,
		IsDeleted: -1,
	})
	if err != nil {
		return 0, err
	}

	return
}
