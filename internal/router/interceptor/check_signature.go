package interceptor

import (
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/constant"
	"net/http"
	"strings"

	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/errors"
	"github.com/xinliangnote/go-gin-api/pkg/signature"
	"github.com/xinliangnote/go-gin-api/pkg/urltable"
)

var whiteListPath = map[string]bool{
	"/login/web": true,
}

func (i *interceptor) CheckSignature() core.HandlerFunc {
	return func(c core.Context) {
		if !env.Active().IsPro() {
			return
		}

		// 签名信息
		authorization := c.GetHeader(configs.HeaderSignToken)
		if authorization == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Authorization 参数")),
			)
			return
		}

		// 时间信息
		date := c.GetHeader(configs.HeaderSignTokenDate)
		if date == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Date 参数")),
			)
			return
		}

		// 通过签名信息获取 key
		authorizationSplit := strings.Split(authorization, " ")
		if len(authorizationSplit) < 2 {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中 Authorization 格式错误")),
			)
			return
		}

		key := authorizationSplit[0]

		data, err := i.authorizedService.DetailByKey(c, key)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(err),
			)
			return
		}

		if data.IsUsed == constant.IsUsedNo {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New(key + " 已被禁止调用")),
			)
			return
		}

		if len(data.Apis) < 1 {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New(key + " 未进行接口授权")),
			)
			return
		}

		if !whiteListPath[c.Path()] {
			// 验证 c.Method() + c.Path() 是否授权
			table := urltable.NewTable()
			for _, v := range data.Apis {
				_ = table.Append(v.Method + v.Api)
			}

			if pattern, _ := table.Mapping(c.Method() + c.Path()); pattern == "" {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.AuthorizationError,
					code.Text(code.AuthorizationError)).WithError(errors.New(c.Method() + c.Path() + " 未进行接口授权")),
				)
				return
			}
		}

		ok, err := signature.New(key, data.Secret, configs.HeaderSignTokenTimeout).Verify(authorization, date, c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(err),
			)
			return
		}

		if !ok {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("Header 中 Authorization 信息错误")),
			)
			return
		}
	}
}
