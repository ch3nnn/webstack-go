package v1

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SSEStream(ctx *gin.Context, dataChan chan interface{}, eventName string) {
	if eventName == "" {
		eventName = "message"
	}

	// 设置响应头
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	// 使用 ctx.Stream 发送数据
	ctx.Stream(func(w io.Writer) bool {
		if data, ok := <-dataChan; ok {
			// 将数据作为 SSE 事件发送
			ctx.SSEvent(eventName, data)
			return true // 继续流
		}
		return false // 关闭流
	})
}

func HandleSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}

	ctx.JSON(http.StatusOK, data)
}

func HandleError(ctx *gin.Context, httpCode int, err error, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	resp := Response{Code: errorCodeMap[err], Message: err.Error(), Data: data}
	if _, ok := errorCodeMap[ErrSuccess]; !ok {
		resp = Response{Code: 500, Message: "unknown error", Data: data}
	}
	ctx.JSON(httpCode, resp)
}

func ErrHandler404(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{"title": "404 Error - Page not found"})
}

type Error struct {
	Code    int
	Message string
}

var errorCodeMap = map[error]int{}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}

func (e Error) Error() string {
	return e.Message
}
