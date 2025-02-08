/**
 * @Author: chentong
 * @Date: 2024/05/27 上午10:23
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	s "github.com/ch3nnn/webstack-go/internal/service"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	// Update 更新分类
	Update(ctx context.Context, req *v1.CategoryUpdateReq) (resp *v1.CategoryUpdateResp, err error)
	// Detail 获取分类详情
	Detail(ctx context.Context, req *v1.CategoryDetailReq) (resp *v1.CategoryDetailResp, err error)
	// List 获取分类列表
	List(ctx context.Context, req *v1.CategoryListReq) (resp *v1.CategoryListResp, err error)
	// Create 创建分类
	Create(ctx context.Context, req *v1.CategoryCreateReq) (resp *v1.CategoryCreateResp, err error)
	// Delete 删除分类
	Delete(ctx context.Context, req *v1.CategoryDeleteReq) (resp *v1.CategoryDeleteResp, err error)
}

type service struct {
	*s.Service
	categoryRepo repository.IStCategoryDao
}

func NewService(
	s *s.Service,
	categoryRepo repository.IStCategoryDao,
) Service {
	return &service{
		Service:      s,
		categoryRepo: categoryRepo,
	}
}

func (s *service) i() {}
