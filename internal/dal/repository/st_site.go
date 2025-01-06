package repository

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

var _ iCustomGenStSiteFunc = (*customStSiteDao)(nil)

type (
	// IStSiteDao not edit interface name
	IStSiteDao interface {
		iWhereStSiteFunc
		WithContext(ctx context.Context) iCustomGenStSiteFunc

		// TODO Custom WhereFunc ....
		// ...
		LikeInByTitleOrDescOrURL(search string) func(dao gen.Dao) gen.Dao
	}

	// not edit interface name
	iCustomGenStSiteFunc interface {
		iGenStSiteFunc

		// TODO Custom DaoFunc ....
		// ...

		FindSiteCategoryWithPage(page, pageSize int, result any, whereFunc ...func(dao gen.Dao) gen.Dao) (count int64, err error)
	}

	// not edit interface name
	customStSiteDao struct {
		stSiteDao
	}
)

type (
	SiteCategory struct {
		model.StSite
		model.StCategory
	}
)

func NewStSiteDao() IStSiteDao {
	return &customStSiteDao{
		stSiteDao{
			stSiteDo: query.StSite.WithContext(context.Background()),
		},
	}
}

func (d *customStSiteDao) WithContext(ctx context.Context) iCustomGenStSiteFunc {
	d.stSiteDo = d.stSiteDo.WithContext(ctx)
	return d
}

func (d *customStSiteDao) LikeInByTitleOrDescOrURL(search string) func(dao gen.Dao) gen.Dao {
	return func(dao gen.Dao) gen.Dao {
		return dao.
			Where(query.StSite.Title.Like("%" + search + "%")).
			Or(query.StSite.Description.Like("%" + search + "%")).
			Or(query.StSite.URL.Like("%" + search + "%"))
	}
}

func (d *customStSiteDao) FindSiteCategoryWithPage(page, pageSize int, result any, whereFunc ...func(dao gen.Dao) gen.Dao) (count int64, err error) {
	return d.stSiteDo.
		Select(
			field.NewField(query.StSite.TableName(), "*"),
			field.NewField(query.StCategory.TableName(), "*"),
		).
		LeftJoin(
			query.StCategory,
			query.StCategory.ID.EqCol(query.StSite.CategoryID),
		).
		Order(query.StSite.CreatedAt.Desc()).
		Scopes(whereFunc...).
		ScanByPage(result, (page-1)*pageSize, pageSize)
}
