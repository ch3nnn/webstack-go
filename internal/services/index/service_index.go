package index

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

func (s *service) Index(ctx core.Context) (sites []*model.Site, err error) {
	sites, err = query.Site.WithContext(ctx.RequestContext()).Find()
	if err != nil {
		return nil, err
	}
	return
}
