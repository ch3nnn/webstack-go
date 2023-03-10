package category

import (
	"net/http"
	"strconv"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type deleteRequest struct {
	Id string `uri:"id"` // 主键 id
}

type deleteResponse struct {
	Id int32 `json:"id"` // 主键ID
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

		id, err := strconv.Atoi(req.Id)
		if err != nil {
			return
		}

		err = h.categoryService.Delete(c, int32(id))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryDeleteError,
				code.Text(code.CategoryDeleteError)).WithError(err),
			)
			return
		}

		res.Id = int32(id)
		c.Payload(res)
	}
}
