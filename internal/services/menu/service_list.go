package menu

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/model"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
)

type SearchData struct {
	Pid int64 // 父类ID
}

func (s *service) List(ctx core.Context, searchData *SearchData) (menus []*model.Menu, err error) {
	iMenuDo := query.Menu.WithContext(ctx.RequestContext()).Where(query.Menu.IsDeleted.Eq(-1))
	if searchData.Pid != 0 {
		iMenuDo = iMenuDo.Where(query.Menu.Pid.Eq(searchData.Pid))
	}
	menus, err = iMenuDo.Order(query.Menu.Sort).Find()
	if err != nil {
		return nil, err
	}

	return
}
