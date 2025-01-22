package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

const (
	DefaultUname     = "admin"
	DefaultUPassword = "admin"
)

var _ iCustomGenSysUserFunc = (*customSysUserDao)(nil)

type (
	// ISysUserDao not edit interface name
	ISysUserDao interface {
		iWhereSysUserFunc
		WithContext(ctx context.Context) iCustomGenSysUserFunc

		// TODO Custom WhereFunc ....
		// ...
	}

	// not edit interface name
	iCustomGenSysUserFunc interface {
		iGenSysUserFunc

		// TODO Custom DaoFunc ....
		// ...
	}

	// not edit interface name
	customSysUserDao struct {
		sysUserDao
	}
)

func NewSysUserDao() ISysUserDao {
	return &customSysUserDao{
		sysUserDao{
			sysUserDo: query.SysUser.WithContext(context.Background()),
		},
	}
}

func (d *customSysUserDao) WithContext(ctx context.Context) iCustomGenSysUserFunc {
	d.sysUserDo = d.sysUserDo.WithContext(ctx)
	return d
}
