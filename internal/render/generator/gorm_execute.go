package generator_handler

import (
	"fmt"
	"github.com/ch3nnn/webstack-go/cmd/gormgen"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"strings"
)

type gormExecuteRequest struct {
	Tables string `form:"tables"`
}

func (h *handler) GormExecute() core.HandlerFunc {

	return func(c core.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Payload(fmt.Sprintf("创建失败! %s", err))
			}
		}()
		req := new(gormExecuteRequest)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.Payload("参数传递有误")
			return
		}
		gormgen.GenerateTable(strings.Split(req.Tables, ","))
		c.Payload("创建成功!")
	}
}
