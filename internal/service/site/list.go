/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:58
 */

package site

import (
	"context"
	"time"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
)

func (s *service) List(ctx context.Context, req *v1.SiteListReq) (resp *v1.SiteListResp, err error) {
	var siteCategories []repository.SiteCategory
	count, err := s.siteRepository.WithContext(ctx).FindSiteCategoryWithPage(req.Page, req.PageSize, &siteCategories, s.siteRepository.LikeInByTitleOrDescOrURL(req.Search))
	if err != nil {
		return nil, err
	}

	list := make([]v1.Site, len(siteCategories))
	for i, siteCategory := range siteCategories {
		list[i] = v1.Site{
			Id:          siteCategory.StSite.ID,
			Thumb:       siteCategory.StSite.Icon,
			Title:       siteCategory.StSite.Title,
			Url:         siteCategory.StSite.URL,
			Category:    siteCategory.StCategory.Title,
			CategoryId:  siteCategory.StSite.CategoryID,
			Description: siteCategory.StSite.Description,
			IsUsed:      siteCategory.StSite.IsUsed,
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
