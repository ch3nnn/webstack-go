/**
 * @Author: chentong
 * @Date: 2024/06/27 下午11:55
 */

package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

// Create
// Register godoc
// @Summary 新增分类
// @Schemes
// @Description 新增分类
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param request body v1.CategoryCreateReq true "params"
// @Success 200 {object} v1.CategoryCreateResp
// @Router /api/admin/category [post]
func (h *Handler) Create(ctx *gin.Context) {
	var req v1.CategoryCreateReq
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	resp, err := h.categoryService.Create(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
