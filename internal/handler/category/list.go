/**
 * @Author: chentong
 * @Date: 2024/06/11 下午11:49
 */

package category

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

// List
// Register godoc
// @Summary 列表分类
// @Schemes
// @Description 列表分类
// @Tags 分类模块
// @Accept json
// @Produce json
// @Success 200 {object} v1.CategoryListResp
// @Router /api/admin/category [get]
func (h *Handler) List(ctx *gin.Context) {
	resp, err := h.categoryService.List(ctx, nil)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
