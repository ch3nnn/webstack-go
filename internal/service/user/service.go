/**
 * @Author: chentong
 * @Date: 2024/05/26 上午12:26
 */

package user

import (
	"context"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	// Info 获取用户信息
	Info(ctx *gin.Context, req *v1.InfoReq) (*v1.InfoResp, error)
	// Login 登录
	Login(ctx context.Context, req *v1.LoginReq) (resp *v1.LoginResp, err error)
	// UpdatePassword 修改密码
	UpdatePassword(ctx *gin.Context, req *v1.UpdatePasswordReq) (*v1.UpdatePasswordResp, error)
}

type service struct {
	*s.Service
	userRepo      repository.ISysUserDao
	siteRepo      repository.IStSiteDao
	categoryRepo  repository.IStCategoryDao
	menuRepo      repository.ISysMenuDao
	adminMenuRepo repository.ISysUserMenuDao
}

func NewService(s *s.Service) Service {
	return &service{
		Service:       s,
		userRepo:      repository.NewSysUserDao(),
		siteRepo:      repository.NewStSiteDao(),
		categoryRepo:  repository.NewStCategoryDao(),
		menuRepo:      repository.NewSysMenuDao(),
		adminMenuRepo: repository.NewSysUserMenuDao(),
	}
}

func (s *service) i() {}
