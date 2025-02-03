package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

// DefaultSysUserMenuAdmin is used to store the menu information of the admin user DefaultSysMenuAdmin
var DefaultSysUserMenuAdmin = []*model.SysUserMenu{
	{
		UserID: 1,
		MenuID: 1,
	},
	{
		UserID: 1,
		MenuID: 2,
	},
	{
		UserID: 1,
		MenuID: 3,
	},
	{
		UserID: 1,
		MenuID: 4,
	},
	{
		UserID: 1,
		MenuID: 5,
	},
}

var _ iCustomGenSysUserMenuFunc = (*customSysUserMenuDao)(nil)

type (
	// ISysUserMenuDao not edit interface name
	ISysUserMenuDao interface {
		iWhereSysUserMenuFunc
		WithContext(ctx context.Context) iCustomGenSysUserMenuFunc

		// TODO Custom WhereFunc ....
		// ...
	}

	// not edit interface name
	iCustomGenSysUserMenuFunc interface {
		iGenSysUserMenuFunc

		// TODO Custom DaoFunc ....
		// ...
	}

	// not edit interface name
	customSysUserMenuDao struct {
		sysUserMenuDao
	}
)

func NewSysUserMenuDao() ISysUserMenuDao {
	return &customSysUserMenuDao{
		sysUserMenuDao{
			sysUserMenuDo: query.SysUserMenu.WithContext(context.Background()),
		},
	}
}

func (d *customSysUserMenuDao) WithContext(ctx context.Context) iCustomGenSysUserMenuFunc {
	d.sysUserMenuDo = d.sysUserMenuDo.WithContext(ctx)
	return d
}
