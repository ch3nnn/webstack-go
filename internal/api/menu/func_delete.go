package menu

import (
	"net/http"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

type deleteRequest struct {
	Id string `uri:"id"` // HashID
}

type deleteResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Tags API.menu
// @Accept json
// @Produce json
// @Param id path string true "hashId"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu/{id} [delete]
// @Security LoginToken
func (h *handler) Delete() core.HandlerFunc {
	return func(c core.Context) {
		req := new(deleteRequest)
		res := new(deleteResponse)
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

		err = h.menuService.Delete(c, int64(ids[0]))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuDeleteError,
				code.Text(code.MenuDeleteError)).WithError(err),
			)
			return
		}

		res.Id = int64(ids[0])
		c.Payload(res)
	}
}
