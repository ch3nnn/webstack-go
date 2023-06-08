package interceptor

import (
	"encoding/json"
	"net/http"

	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/errors"
	"github.com/ch3nnn/webstack-go/internal/pkg/urltable"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
	"github.com/ch3nnn/webstack-go/internal/services/admin"
)

func (i *interceptor) CheckRBAC() core.HandlerFunc {
	return func(c core.Context) {
		token := c.GetHeader("Token")
		if token == "" {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Token 参数")),
			)
			return
		}

		if !i.cache.Exists(configs.RedisKeyPrefixLoginUser + token) {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithError(errors.New("请先登录")),
			)
			return
		}

		if !i.cache.Exists(configs.RedisKeyPrefixLoginUser + token + ":action") {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithError(errors.New("当前账号未配置 RBAC 权限")),
			)
			return
		}

		actionData, err := i.cache.Get(configs.RedisKeyPrefixLoginUser+token+":action", redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.CacheGetError,
				code.Text(code.CacheGetError)).WithError(err),
			)
			return
		}

		var actions []admin.MyActionData
		err = json.Unmarshal([]byte(actionData), &actions)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(err),
			)
			return
		}

		if len(actions) > 0 {
			table := urltable.NewTable()
			for _, v := range actions {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.RBACError,
					code.Text(code.RBACError)).WithError(errors.New(c.Method() + c.Path() + " 未进行 RBAC 授权")),
				)
				return
			}
		}

	}
}
