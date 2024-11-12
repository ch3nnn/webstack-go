/**
 * @Author: chentong
 * @Date: 2024/06/28 上午12:25
 */

package site

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Delete(ctx *gin.Context) {
	var req v1.SiteDeleteReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	resp, err := h.siteService.Delete(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
