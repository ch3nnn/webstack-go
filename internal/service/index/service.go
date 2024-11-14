/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:48
 */

package index

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Index(ctx context.Context) (*v1.IndexResponseData, error)
}

type service struct {
	siteRepo     repository.IStSiteDao
	categoryRepo repository.IStCategoryDao
	*s.Service
}

func NewService(s *s.Service) Service {
	return &service{
		Service:      s,
		siteRepo:     repository.NewStSiteDao(),
		categoryRepo: repository.NewStCategoryDao(),
	}
}

func (s *service) i() {}
