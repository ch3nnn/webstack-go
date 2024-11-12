/**
 * @Author: chentong
 * @Date: 2024/05/26 下午3:51
 */

package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/internal/middleware"
)

func (s *service) Info(ctx *gin.Context, _ *v1.InfoReq) (*v1.InfoResp, error) {
	var (
		g          errgroup.Group
		user       *model.SysUser
		menus      []*model.SysMenu
		adminMenus []*model.SysUserMenu
	)

	g.Go(func() (err error) {
		user, err = s.userRepo.WithContext(ctx).FindOne(s.userRepo.WhereByID(ctx.GetInt(middleware.UserID)))
		if err != nil {
			return err
		}
		return nil
	})

	g.Go(func() (err error) {
		menus, err = s.menuRepo.WithContext(ctx).FindAll()
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() (err error) {
		adminMenus, err = s.adminMenuRepo.WithContext(ctx).FindAll(s.adminMenuRepo.WhereByUserID(ctx.GetInt(middleware.UserID)))
		if err != nil {
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	var menuList []v1.Menu
	for _, menu := range menus {
		for _, adminMenu := range adminMenus {
			if menu.ID == adminMenu.MenuID {
				menuList = append(menuList, v1.Menu{
					Id:   menu.ID,
					Pid:  menu.Pid,
					Name: menu.Name,
					Link: menu.Link,
					Icon: menu.Icon,
				})
			}
		}
	}

	return &v1.InfoResp{
		Username: user.Username,
		Menus:    menuList,
	}, nil
}
