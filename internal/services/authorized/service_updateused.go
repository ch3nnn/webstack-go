package authorized

import (
	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"

	"gorm.io/gorm"
)

func (s *service) UpdateUsed(ctx core.Context, id, used int64) (err error) {
	authorizedInfo, err := query.Authorized.WithContext(ctx.RequestContext()).
		Where(query.Authorized.IsDeleted.Eq(-1)).
		Where(query.Authorized.ID.Eq(id)).
		First()
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	_, err = query.Authorized.WithContext(ctx.RequestContext()).
		Where(query.Authorized.ID.Eq(id)).
		UpdateColumnSimple(
			query.Authorized.IsUsed.Value(used),
			query.Authorized.UpdatedUser.Value(ctx.SessionUserInfo().UserName),
		)

	if err != nil {
		return err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedInfo.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}
