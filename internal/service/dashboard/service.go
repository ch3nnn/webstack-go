/**
 * @Author: chentong
 * @Date: 2025/01/17 下午7:32
 */

package dashboard

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

	Dashboard(ctx *gin.Context) (*v1.DashboardResp, error)

	// Stat 统计卡片与近 7 天趋势
	Stat(ctx context.Context) (*v1.StatResp, error)
}

type service struct {
	*s.Service
	siteRepo  repository.IStSiteDao
	visitRepo repository.ISysVisitDao
}

func NewService(s *s.Service, siteRepo repository.IStSiteDao, visitRepo repository.ISysVisitDao) Service {
	return &service{
		Service:   s,
		siteRepo:  siteRepo,
		visitRepo: visitRepo,
	}
}

func (s *service) i() {}
