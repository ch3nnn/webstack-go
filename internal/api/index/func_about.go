package index

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
)

func (h *handler) About() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("about", nil)
	}

}
