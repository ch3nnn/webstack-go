/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:46
 */

package category

import (
	"github.com/ch3nnn/webstack-go/internal/handler"
	"github.com/ch3nnn/webstack-go/internal/service/category"
)

type Handler struct {
	*handler.Handler
	categoryService category.Service
}

func NewHandler(handler *handler.Handler, categoryService category.Service) *Handler {
	return &Handler{
		Handler:         handler,
		categoryService: categoryService,
	}
}
