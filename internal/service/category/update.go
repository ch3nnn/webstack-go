/**
 * @Author: chentong
 * @Date: 2024/06/13 下午11:17
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/pkg/gormx"
)

func (s *service) Update(ctx context.Context, req *v1.CategoryUpdateReq) (*v1.CategoryUpdateResp, error) {
	update := make(map[string]any)

	if req.Pid != nil {
		column := gormx.ColumnName(query.StCategory.ParentID)
		update[column] = req.Pid
	}
	if req.Icon != nil {
		column := gormx.ColumnName(query.StCategory.Icon)
		update[column] = req.Icon
	}
	if req.Name != nil {
		column := gormx.ColumnName(query.StCategory.Title)
		update[column] = req.Name
	}
	if req.SortID != nil {
		column := gormx.ColumnName(query.StCategory.Sort)
		update[column] = req.SortID
	}
	if req.IsUsed != nil {
		column := gormx.ColumnName(query.StCategory.IsUsed)
		update[column] = req.IsUsed
	}

	_, err := s.categoryRepo.WithContext(ctx).Update(update, s.categoryRepo.WhereByID(req.ID))
	if err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.WithContext(ctx).FindOne(s.categoryRepo.WhereByID(req.ID))
	if err != nil {
		return nil, err
	}

	return &v1.CategoryUpdateResp{
		Category: v1.Category{
			ID:        category.ID,
			ParentID:  category.ParentID,
			Sort:      category.Sort,
			Title:     category.Title,
			Icon:      category.Icon,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			IsUsed:    category.IsUsed,
			Level:     category.Level,
		},
	}, nil
}
