/**
 * @Author: chentong
 * @Date: 2025/01/17 下午7:32
 */

package dashboard

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	s "github.com/ch3nnn/webstack-go/internal/service"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	Dashboard(ctx *gin.Context) (*v1.DashboardResp, error)
}

type service struct {
	*s.Service
}

func NewService(s *s.Service) Service {
	return &service{
		Service: s,
	}
}

func (s *service) i() {}
