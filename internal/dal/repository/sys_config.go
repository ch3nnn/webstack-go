package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

var _ iCustomGenSysConfigFunc = (*customSysConfigDao)(nil)

type (
	// ISysConfigDao not edit interface name
	ISysConfigDao interface {
		iWhereSysConfigFunc
		WithContext(ctx context.Context) iCustomGenSysConfigFunc

		// TODO Custom WhereFunc ....
		// ...
	}

	// not edit interface name
	iCustomGenSysConfigFunc interface {
		iGenSysConfigFunc

		// TODO Custom DaoFunc ....
		// ...
	}

	// not edit interface name
	customSysConfigDao struct {
		sysConfigDao
	}
)

func NewSysConfigDao() ISysConfigDao {
	return &customSysConfigDao{
		sysConfigDao{
			sysConfigDo: query.SysConfig.WithContext(context.Background()),
		},
	}
}

func (d *customSysConfigDao) WithContext(ctx context.Context) iCustomGenSysConfigFunc {
	d.sysConfigDo = d.sysConfigDo.WithContext(ctx)
	return d
}
