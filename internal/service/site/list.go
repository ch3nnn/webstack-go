/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:58
 */

package site

import (
	"context"
	"time"

	"gorm.io/gen"
	"gorm.io/gen/field"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
)

func (s *service) List(ctx context.Context, req *v1.SiteListReq) (resp *v1.SiteListResp, err error) {
	var orderColumns []field.Expr
	orderColumns = append(orderColumns, query.StSite.CreatedAt.Desc())

	var whereFunc []func(dao gen.Dao) gen.Dao
	if req.Search != "" {
		whereFunc = append(whereFunc, s.siteRepository.LikeInByTitleOrDescOrURL(req.Search))
	}
	if req.CategoryID != 0 {
		whereFunc = append(whereFunc, s.siteRepository.WhereByCategoryID(req.CategoryID))
		orderColumns = []field.Expr{query.StSite.Sort.Asc()} // 同分类网址按排序升序
	}

	var siteCategories []repository.SiteCategory
	count, err := s.siteRepository.WithContext(ctx).FindSiteCategoryWithPage(req.Page, req.PageSize, &siteCategories, orderColumns, whereFunc...)
	if err != nil {
		return nil, err
	}

	list := make([]v1.Site, len(siteCategories))
	for i, siteCategory := range siteCategories {
		list[i] = v1.Site{
			Id:          siteCategory.StSite.ID,
			Icon:        siteCategory.StSite.Icon,
			Title:       siteCategory.StSite.Title,
			Url:         siteCategory.StSite.URL,
			Category:    siteCategory.StCategory.Title,
			CategoryId:  siteCategory.StSite.CategoryID,
			Description: siteCategory.StSite.Description,
			IsUsed:      siteCategory.StSite.IsUsed,
			Sort:        siteCategory.StSite.Sort,
			CreatedAt:   siteCategory.StSite.CreatedAt.Format(time.DateTime),
			UpdatedAt:   siteCategory.StSite.UpdatedAt.Format(time.DateTime),
		}
	}

	return &v1.SiteListResp{
		List: list,
		Pagination: v1.SiteLisPagination{
			Total:        count,
			CurrentPage:  req.Page,
			PerPageCount: req.PageSize,
		},
	}, err
}
