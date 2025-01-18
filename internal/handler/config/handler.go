/**
 * @Author: chentong
 * @Date: 2025/01/17 下午7:32
 */

package config

import (
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/config"
)

type Handler struct {
	*handler.Handler
	configService config.Service
}

func NewHandler(handler *handler.Handler, configService config.Service) *Handler {
	return &Handler{
		Handler:       handler,
		configService: configService,
	}
}
