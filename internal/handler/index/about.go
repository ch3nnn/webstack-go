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

// AboutPage 渲染关于页 HTML：导航栏要用的 logo / 标题来自站点配置，
// 正文仍由页面内的 /api/about 拉取后渲染 markdown。
func (h *Handler) AboutPage(ctx *gin.Context) {
	resp, err := h.indexService.About(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	ctx.HTML(http.StatusOK, "about.html", resp)
}
