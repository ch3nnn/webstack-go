package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/ch3nnn/webstack-go/pkg/jwt"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

// IndexAuth 首页认证中间件 - 只检查首页访问权限
func IndexAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查是否有有效的JWT token
		tokenString := ctx.Request.Header.Get("Token")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("webstack_session")
		}
		if tokenString == "" {
			tokenString = ctx.Query("Token")
			if tokenString == "" {
				tokenString = ctx.Query("token")
			}
		}

		// 如果没有token，重定向到飞书登录
		if tokenString == "" {
			logger.WithContext(ctx).Info("No token found, redirecting to feishu login")
			ctx.Redirect(http.StatusFound, "/api/auth/feishu/login")
			ctx.Abort()
			return
		}

		// 验证token
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			logger.WithContext(ctx).Error("Invalid token", zap.Error(err))
			ctx.Redirect(http.StatusFound, "/api/auth/feishu/login")
			ctx.Abort()
			return
		}

		// 设置用户信息到上下文
		ctx.Set(UserID, claims.UserID)
		ctx.Set(Claims, claims)

		logger.WithContext(ctx).Info("User authenticated", zap.Int("UserID", claims.UserID))
		ctx.Next()
	}
}
