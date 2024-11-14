/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package index

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Index(ctx *gin.Context) {
	resp, err := h.indexService.Index(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.HTML(http.StatusOK, "index.html", resp)
}
