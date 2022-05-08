package routes

import (
	"github.com/gin-gonic/gin"
	"go-web-site-admin/src/go/controller"
)

func SetUserRouter() *gin.Engine {
	router := gin.Default()

	routerGroup := router.Group("api/v1/user")
	{
		// 获取用户信息
		routerGroup.GET("/users", controller.UserList)
	}
	return router
}
