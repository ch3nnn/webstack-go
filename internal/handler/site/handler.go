/**
 * @Author: chentong
 * @Date: 2024/05/26 上午12:35
 */

package site

import (
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/site"
)

type Handler struct {
	*handler.Handler
	siteService site.Service
}

func NewHandler(handler *handler.Handler, siteService site.Service) *Handler {
	return &Handler{
		Handler:     handler,
		siteService: siteService,
	}
}
