package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

var _ iCustomGenSysMenuFunc = (*customSysMenuDao)(nil)

type (
	// ISysMenuDao not edit interface name
	ISysMenuDao interface {
		iWhereSysMenuFunc
		WithContext(ctx context.Context) iCustomGenSysMenuFunc

		// TODO Custom WhereFunc ....
		// ...
	}

	// not edit interface name
	iCustomGenSysMenuFunc interface {
		iGenSysMenuFunc

		// TODO Custom DaoFunc ....
		// ...
	}

	// not edit interface name
	customSysMenuDao struct {
		sysMenuDao
	}
)

func NewSysMenuDao() ISysMenuDao {
	return &customSysMenuDao{
		sysMenuDao{
			sysMenuDo: query.SysMenu.WithContext(context.Background()),
		},
	}
}

func (d *customSysMenuDao) WithContext(ctx context.Context) iCustomGenSysMenuFunc {
	d.sysMenuDo = d.sysMenuDo.WithContext(ctx)
	return d
}
