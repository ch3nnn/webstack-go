/**
 * @Author: chentong
 * @Date: 2024/06/27 下午11:30
 */

package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

// Delete
// Register godoc
// @Summary 删除分类
// @Schemes
// @Description 删除分类
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param request body v1.CategoryDeleteReq true "params"
// @Success 200 {object} v1.CategoryDeleteResp
// @Router /api/admin/category/:id [delete]
func (h *Handler) Delete(ctx *gin.Context) {
	var req v1.CategoryDeleteReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	if _, err := h.categoryService.Delete(ctx, &req); err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, nil)
}
