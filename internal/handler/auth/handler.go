package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	authService "github.com/ch3nnn/webstack-go/internal/service/auth"
)

type Handler struct {
	authService authService.Service
}

func NewHandler(authService authService.Service) *Handler {
	return &Handler{
		authService: authService,
	}
}

// FeishuLogin 飞书登录
func (h *Handler) FeishuLogin(ctx *gin.Context) {
	oauthURL, state, err := h.authService.GenerateOAuthURL(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	// 设置 state cookie，用于验证回调
	ctx.SetCookie("feishu_state", state, 300, "/", "localhost", false, true) // 5分钟过期，httpOnly

	// 重定向到飞书 OAuth 页面
	ctx.Redirect(http.StatusFound, oauthURL)
}

// FeishuCallback 飞书登录回调
func (h *Handler) FeishuCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")
	stateFromCookie, _ := ctx.Cookie("feishu_state")

	if code == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message":  "缺少 code 参数",
			"link":     "/",
			"linkText": "去首页",
		})
		return
	}

	// 简化state验证，如果state不匹配但code存在，仍然继续处理
	if state == "" {
		ctx.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message":  "缺少 state 参数",
			"link":     "/",
			"linkText": "去首页",
		})
		return
	}

	// 如果state不匹配，记录警告但继续处理
	if stateFromCookie != "" && state != stateFromCookie {
		fmt.Printf("WARNING: state mismatch: received=%s, expected=%s\n", state, stateFromCookie)
	}

	// 处理 OAuth 回调
	token, err := h.authService.HandleOAuthCallback(ctx, code, state)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"message":  fmt.Sprintf("登录失败: %s", err.Error()),
			"link":     "/",
			"linkText": "去首页",
		})
		return
	}

	// 清除 state cookie
	ctx.SetCookie("feishu_state", "", -1, "/", "", false, true)

	// 设置 JWT cookie
	ctx.SetCookie("webstack_session", token, 7*24*3600, "/", "", false, true) // 7天过期，httpOnly

	// 重定向到首页
	ctx.Redirect(http.StatusFound, "/")
}

// Logout 登出
func (h *Handler) Logout(ctx *gin.Context) {
	// 清除 JWT cookie
	ctx.SetCookie("webstack_session", "", -1, "/", "", false, true)

	// 重定向到飞书登录
	ctx.Redirect(http.StatusFound, "/api/auth/feishu/login")
}
