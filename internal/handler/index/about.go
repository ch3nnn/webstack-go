/**
 * @Author: chentong
 * @Date: 2025/01/18 20:45
 */

package index

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) About(ctx *gin.Context) {
	resp, err := h.indexService.About(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
