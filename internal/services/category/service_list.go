package category

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Pid int32 // 父类ID
}

func (s *service) List(ctx core.Context) (categories []*model.Category, err error) {
	categories, err = query.Category.WithContext(ctx.RequestContext()).
		Order(query.Category.Sort).
		Find()
	if err != nil {
		return nil, err
	}

	return
}
