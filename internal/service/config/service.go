/**
 * @Author: chentong
 * @Date: 2025/01/17 下午7:32
 */

package config

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

	// GetConfig 获取配置信息
	GetConfig(ctx context.Context) (*v1.ConfigResp, error)
	// Update 更新配置信息
	Update(ctx *gin.Context, req *v1.ConfigUpdateReq) (resp *v1.ConfigUpdateResp, err error)
}

type service struct {
	*s.Service
	configRepo repository.ISysConfigDao
}

func NewService(
	s *s.Service,
	configRepo repository.ISysConfigDao,
) Service {
	return &service{
		Service:    s,
		configRepo: configRepo,
	}
}

func (s *service) i() {}
