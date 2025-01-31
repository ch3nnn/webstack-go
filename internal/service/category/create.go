/**
 * @Author: chentong
 * @Date: 2024/05/27 上午11:14
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func (s *service) Create(ctx context.Context, req *v1.CategoryCreateReq) (*v1.CategoryCreateResp, error) {
	category, err := s.categoryRepo.WithContext(ctx).
		Create(&model.StCategory{
			ParentID: req.ParentID,
			Title:    req.Name,
			Icon:     req.Icon,
			Level:    req.Level,
			IsUsed:   req.IsUsed,
			Sort:     req.SortID,
		})
	if err != nil {
		return nil, err
	}

	return &v1.CategoryCreateResp{Category: v1.Category{
		ID:        category.ID,
		ParentID:  category.ParentID,
		Sort:      category.Sort,
		Title:     category.Title,
		Icon:      category.Icon,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		IsUsed:    category.IsUsed,
		Level:     category.Level,
	}}, nil
}
