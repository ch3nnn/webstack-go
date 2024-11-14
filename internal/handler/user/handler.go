/**
 * @Author: chentong
 * @Date: 2024/05/26 上午12:35
 */

package user

import (
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/user"
)

type Handler struct {
	*handler.Handler
	userService user.Service
}

func NewHandler(handler *handler.Handler, userService user.Service) *Handler {
	return &Handler{
		Handler:     handler,
		userService: userService,
	}
}
