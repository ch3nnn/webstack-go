package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-site-admin/src/go/common"
)

//
// UserList
// @Description:
// @param c
//
func UserList(c *gin.Context) {
	common.Success(c, "请求成功", nil)
}
