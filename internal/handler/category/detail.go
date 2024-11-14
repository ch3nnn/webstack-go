/**
 * @Author: chentong
 * @Date: 2024/06/13 下午11:55
 */

package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

// Detail
// Register godoc
// @Summary 详情分类
// @Schemes
// @Description 详情分类
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param request body v1.CategoryDetailReq true "params"
// @Success 200 {object} v1.CategoryDetailResp
// @Router /api/admin/category/:id [get]
func (h *Handler) Detail(ctx *gin.Context) {
	var req v1.CategoryDetailReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	resp, err := h.categoryService.Detail(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
