package admin

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/configs"
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/pkg/password"
	"github.com/ch3nnn/webstack-go/internal/repository/redis"
)

type offlineRequest struct {
	Id string `form:"id"` // 主键ID
}

type offlineResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Offline 下线管理员
// @Summary 下线管理员
// @Description 下线管理员
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "Hashid"
// @Success 200 {object} offlineResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/offline [patch]
// @Security LoginToken
func (h *handler) Offline() core.HandlerFunc {
	return func(c core.Context) {
		req := new(offlineRequest)
		res := new(offlineResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int64(ids[0])

		b := h.cache.Del(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(id), redis.WithTrace(c.Trace()))
		if !b {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminOfflineError,
				code.Text(code.AdminOfflineError)),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
