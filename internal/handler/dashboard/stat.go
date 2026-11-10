package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

// Stat 仪表盘统计：链接数量、今日 PV/UV、累计 PV 与近 7 天趋势
func (h *Handler) Stat(ctx *gin.Context) {
	resp, err := h.dashboardService.Stat(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	v1.HandleSuccess(ctx, resp)
}
