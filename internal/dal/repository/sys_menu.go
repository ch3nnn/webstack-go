package repository

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
)

var DefaultSysMenuAdmin = []*model.SysMenu{
	{
		ID:     1,
		Pid:    0,
		Name:   "网站管理",
		Icon:   "users",
		Level:  1,
		Sort:   500,
		IsUsed: true,
	},
	{
		ID:     2,
		Pid:    1,
		Name:   "网站分类",
		Link:   "/admin/category",
		Level:  1,
		Sort:   501,
		IsUsed: true,
	},
	{
		ID:     3,
		Pid:    1,
		Name:   "网站信息",
		Link:   "/admin/site",
		Level:  1,
		Sort:   502,
		IsUsed: true,
	},
	{
		ID:     4,
		Pid:    0,
		Name:   "系统管理",
		Level:  1,
		Sort:   600,
		IsUsed: true,
	},
	{
		ID:     5,
		Pid:    4,
		Name:   "网站配置",
		Link:   "/admin/config",
		Level:  1,
		Sort:   601,
		IsUsed: true,
	},
}

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
