package generator_handler

import (
	"github.com/ch3nnn/webstack-go/cmd/handlergen"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

type handlerExecuteRequest struct {
	Name string `form:"name"`
}

func (h *handler) HandlerExecute() core.HandlerFunc {

	return func(c core.Context) {
		req := new(handlerExecuteRequest)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.Payload("参数传递有误")
			return
		}
		if err := handlergen.GenerateHandler(req.Name); err != nil {
			c.Payload(err.Error())
			return
		}
		c.Payload("创建成功!")
	}
}
