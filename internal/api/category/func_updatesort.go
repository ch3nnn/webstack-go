package category

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"net/http"
)

type updateSortRequest struct {
	Id   int64 `form:"id"`   // 主键 id
	Sort int64 `form:"sort"` // 排序
}

type updateSortResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// UpdateSort 更新分类排序
// @Summary 更新分类排序
// @Description 更新分类排序
// @Tags API.category
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashId"
// @Param sort formData int true "排序"
// @Success 200 {object} updateSortResponse
// @Failure 400 {object} code.Failure
// @Router /api/category/sort [patch]
// @Security LoginToken
func (h *handler) UpdateSort() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateSortRequest)
		res := new(updateSortResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		if err := h.categoryService.UpdateSort(c, req.Id, req.Sort); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryUpdateError,
				code.Text(code.CategoryUpdateError)).WithError(err),
			)
			return
		}

		res.Id = req.Id
		c.Payload(res)
	}
}
