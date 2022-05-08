package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultCode struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//
// Success
//  @Description: 请求成功
//  @param c
//  @param msg
//  @param data
//
func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
		"data": data,
	})
}

//
// Error
//  @Description: 请求失败
//  @param c
//  @param msg
//
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 0,
		"msg":  msg,
		"data": nil,
	})

}
