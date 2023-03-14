package category

import (
	"net/http"
	"strconv"

	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

type updateSortRequest struct {
	Id   string `form:"id"`   // 主键 id
	Sort int32  `form:"sort"` // 排序
}

type updateSortResponse struct {
	Id int32 `json:"id"` // 主键ID
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
		id, err := strconv.Atoi(req.Id)
		if err != nil {
			return
		}

		err = h.categoryService.UpdateSort(c, int32(id), req.Sort)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CategoryUpdateError,
				code.Text(code.CategoryUpdateError)).WithError(err),
			)
			return
		}

		res.Id = int32(id)
		c.Payload(res)
	}
}
