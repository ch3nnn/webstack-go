/**
 * @Author: chentong
 * @Date: 2025/01/17 21:21
 */

package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Config(ctx *gin.Context) {
	resp, err := h.configService.GetConfig(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
