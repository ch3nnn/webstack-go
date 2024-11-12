/**
 * @Author: chentong
 * @Date: 2024/05/26 下午4:00
 */

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Info(ctx *gin.Context) {
	resp, err := h.userService.Info(ctx, nil)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
