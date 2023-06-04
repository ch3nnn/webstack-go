package authorized

import (
	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

type CreateAuthorizedAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 请求方法
	API         string `json:"api"`          // 请求地址
}

func (s *service) CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int64, err error) {
	err = query.AuthorizedAPI.WithContext(ctx.RequestContext()).Create(&model.AuthorizedAPI{
		BusinessKey: authorizedAPIData.BusinessKey,
		Method:      authorizedAPIData.Method,
		API:         authorizedAPIData.API,
		IsDeleted:   -1,
		CreatedUser: ctx.SessionUserInfo().UserName,
	})
	if err != nil {
		return 0, err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedAPIData.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}
