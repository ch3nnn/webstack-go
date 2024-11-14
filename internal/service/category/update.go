/**
 * @Author: chentong
 * @Date: 2024/06/13 下午11:17
 */

package category

import (
	"context"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Update(ctx context.Context, req *v1.CategoryUpdateReq) (*v1.CategoryUpdateResp, error) {
	update := make(map[string]any)

	if req.Pid != nil {
		update["parent_id"] = req.Pid
	}
	if req.Icon != nil {
		update["icon"] = req.Icon
	}
	if req.Name != nil {
		update["title"] = req.Name
	}
	if req.Icon != nil {
		update["icon"] = req.Icon
	}
	if req.Sort != nil {
		update["sort"] = req.Sort
	}
	if req.IsUsed != nil {
		update["is_used"] = req.IsUsed
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
