/**
 * @Author: chentong
 * @Date: 2024/06/04 下午5:55
 */

package site

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Create(ctx *gin.Context) {
	var req v1.SiteCreateReq
	if err := ctx.ShouldBind(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	resp, err := h.siteService.BatchCreate(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
