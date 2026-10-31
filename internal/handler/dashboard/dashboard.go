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
	done := ctx.Request.Context().Done()

	go func() {
		defer close(messageChan)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				dashboard, err := h.dashboardService.Dashboard(ctx)
				if err != nil {
					h.Logger.Error("SSE（Server-Sent Events）dashboard api", zap.Error(err))
					return
				}
				// 消费端（SSEStream）停止读取时不能在通道上死等，否则 goroutine 泄漏
				select {
				case messageChan <- dashboard:
				case <-done:
					return
				}
			}
		}
	}()

	v1.SSEStream(ctx, messageChan, "dashboard")
}
