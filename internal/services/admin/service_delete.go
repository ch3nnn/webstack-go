package admin

import (
	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/password"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

func (s *service) Delete(ctx core.Context, id int64) (err error) {
	if _, err = query.Admin.WithContext(ctx.RequestContext()).
		Where(query.Admin.ID.Eq(id)).
		UpdateColumnSimple(query.Admin.IsDeleted.Value(1),
			query.Admin.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		); err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(ctx.Trace()))
	return
}
