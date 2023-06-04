package authorized

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"io"
)

type CreateAuthorizedData struct {
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
}

func (s *service) Create(ctx core.Context, authorizedData *CreateAuthorizedData) (id int64, err error) {
	buf := make([]byte, 10)
	io.ReadFull(rand.Reader, buf)
	secret := hex.EncodeToString(buf)

	err = query.Authorized.WithContext(ctx.RequestContext()).Create(&model.Authorized{
		BusinessKey:       authorizedData.BusinessKey,
		BusinessSecret:    secret,
		BusinessDeveloper: authorizedData.BusinessDeveloper,
		Remark:            authorizedData.Remark,
		IsUsed:            1,
		IsDeleted:         -1,
		CreatedUser:       ctx.SessionUserInfo().UserName,
	})

	if err != nil {
		return 0, err
	}
	return
}
