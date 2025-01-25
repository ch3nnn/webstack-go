/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:48
 */

package index

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

	// Index 首页
	Index(ctx context.Context) (*v1.IndexResp, error)
	// About 关于我
	About(ctx *gin.Context) (*v1.AboutResp, error)
}

type service struct {
	siteRepo     repository.IStSiteDao
	categoryRepo repository.IStCategoryDao
	configRepo   repository.ISysConfigDao
	*s.Service
}

func NewService(s *s.Service) Service {
	return &service{
		Service:      s,
		siteRepo:     repository.NewStSiteDao(),
		categoryRepo: repository.NewStCategoryDao(),
		configRepo:   repository.NewSysConfigDao(),
	}
}

func (s *service) i() {}
