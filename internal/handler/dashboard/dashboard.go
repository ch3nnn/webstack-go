/**
 * @Author: chentong
 * @Date: 2025/02/07 19:48
 */

package dashboard

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (h *Handler) Dashboard(ctx *gin.Context) {
	messageChan := make(chan any)
	go func() {
		for {
			dashboard, err := h.dashboardService.Dashboard(ctx)
			if err != nil {
				h.Logger.Error("SSE（Server-Sent Events）dashboard api", zap.Error(err))
				return
			}

			select {
			case messageChan <- dashboard:
			default:
				return
			}

			time.Sleep(1 * time.Second)
		}
	}()

	v1.SSEStream(ctx, messageChan, "dashboard")
}
