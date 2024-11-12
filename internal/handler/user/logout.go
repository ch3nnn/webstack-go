/**
 * @Author: chentong
 * @Date: 2024/11/12 13:26
 */

package user

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Logout(ctx *gin.Context) {
	v1.HandleSuccess(ctx, nil)
}
