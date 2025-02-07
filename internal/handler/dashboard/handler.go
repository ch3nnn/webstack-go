/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package index

import (
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/dashboard"
)

type Handler struct {
	*handler.Handler
	dashboardService dashboard.Service
}

func NewHandler(handler *handler.Handler, dashboardService dashboard.Service) *Handler {
	return &Handler{Handler: handler, dashboardService: dashboardService}
}
