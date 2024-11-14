/**
 * @Author: chentong
 * @Date: 2024/06/13 下午11:13
 */

package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

// Update
// Register godoc
// @Summary 更新分类
// @Schemes
// @Description 更新分类
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param request body v1.CategoryUpdateReq true "params"
// @Success 200 {object} v1.CategoryUpdateResp
// @Router /api/admin/category/update [put]
func (h *Handler) Update(ctx *gin.Context) {
	var req v1.CategoryUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	resp, err := h.categoryService.Update(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
