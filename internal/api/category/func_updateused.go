package category

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"net/http"
)

type updateUsedRequest struct {
	Id   int64 `form:"id"`   // 主键ID
	Used int64 `form:"used"` // 是否启用 1:是 -1:否
}

type updateUsedResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// UpdateUsed 更新分类为启用/禁用
// @Summary 更新分类为启用/禁用
// @Description 更新分类为启用/禁用
// @Tags API.category
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashId"
// @Param used formData int true "是否启用 1:是 -1:否"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/category/used [patch]
// @Security LoginToken
func (h *handler) UpdateUsed() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		if err := h.categoryService.UpdateUsed(c, req.Id, req.Used); err != nil {
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
