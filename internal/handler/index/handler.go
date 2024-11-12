/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package index

import (
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/index"
)

type Handler struct {
	*handler.Handler
	indexService index.Service
}

func NewHandler(handler *handler.Handler, indexService index.Service) *Handler {
	return &Handler{Handler: handler, indexService: indexService}
}
