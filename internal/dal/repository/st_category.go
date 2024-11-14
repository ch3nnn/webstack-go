package repository

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

var _ iCustomGenStCategoryFunc = (*customStCategoryDao)(nil)

type (
	// IStCategoryDao not edit interface name
	IStCategoryDao interface {
		iWhereStCategoryFunc
		WithContext(ctx context.Context) iCustomGenStCategoryFunc

		// TODO Custom WhereFunc ....
		// ...
	}

	// not edit interface name
	iCustomGenStCategoryFunc interface {
		iGenStCategoryFunc

		// TODO Custom DaoFunc ....
		// ...

		FindAllOrderBySort(orderColumn field.Expr, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.StCategory, error)
	}

	// not edit interface name
	customStCategoryDao struct {
		stCategoryDao
	}
)

func NewStCategoryDao() IStCategoryDao {
	return &customStCategoryDao{
		stCategoryDao{
			stCategoryDo: query.StCategory.WithContext(context.Background()),
		},
	}
}

func (d *customStCategoryDao) WithContext(ctx context.Context) iCustomGenStCategoryFunc {
	d.stCategoryDo = d.stCategoryDo.WithContext(ctx)
	return d
}

func (d *customStCategoryDao) FindAllOrderBySort(orderColumn field.Expr, whereFunc ...func(dao gen.Dao) gen.Dao) ([]*model.StCategory, error) {
	return d.stCategoryDo.Scopes(whereFunc...).Order(orderColumn).Find()
}
