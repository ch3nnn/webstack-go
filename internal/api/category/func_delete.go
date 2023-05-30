package category

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"net/http"
)

type deleteRequest struct {
	Id int64 `uri:"id"` // 主键 id
}

type deleteResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// Delete 删除分类
// @Summary 删除分类
// @Description 删除分类
// @Tags API.category
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} deleteResponse
// @Failure 400 {object} code.Failure
// @Router /api/category/{id} [delete]
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

		if err := h.categoryService.Delete(c, req.Id); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryDeleteError,
				code.Text(code.CategoryDeleteError)).WithError(err),
			)
			return
		}

		res.Id = req.Id
		c.Payload(res)
	}
}
