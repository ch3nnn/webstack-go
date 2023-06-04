package authorized

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

type deleteAPIRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteAPIResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// DeleteAPI 删除调用方接口地址
// @Summary 删除调用方接口地址
// @Description 删除调用方接口地址
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param id path string true "主键ID"
// @Success 200 {object} deleteAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api/{id} [delete]
// @Security LoginToken
func (h *handler) DeleteAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteAPIRequest)
		res := new(deleteAPIResponse)
		if err := c.ShouldBindURI(req); err != nil {
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

		err = h.authorizedService.DeleteAPI(c, id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedDeleteAPIError,
				code.Text(code.AuthorizedDeleteAPIError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
