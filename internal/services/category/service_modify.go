package category

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/category"
)

type UpdateCategoryData struct {
	Name string // 菜单名称
	Icon string // 图标
}

func (s *service) Modify(ctx core.Context, id int32, categoryData *UpdateCategoryData) (err error) {
	data := map[string]interface{}{
		"title": categoryData.Name,
		"icon":  categoryData.Icon,
	}

	qb := category.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, id)
	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	return
}
