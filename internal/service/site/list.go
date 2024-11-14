/**
 * @Author: chentong
 * @Date: 2024/05/27 下午5:58
 */

package site

import (
	"context"
	"time"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) List(ctx context.Context, req *v1.SiteListReq) (resp *v1.SiteListResp, err error) {
	sites, count, err := s.siteRepository.WithContext(ctx).FindPage(req.Page, req.PageSize, nil, s.siteRepository.LikeInByTitleOrDescOrURL(req.Search))
	if err != nil {
		return nil, err
	}

	list := make([]v1.Site, len(sites))
	for i, site := range sites {

		var categoryName string
		category, _ := s.categoryRepository.WithContext(ctx).FindOne(s.categoryRepository.WhereByID(site.CategoryID))
		if category != nil {
			categoryName = category.Title
		}

		list[i] = v1.Site{
			Id:          site.ID,
			Thumb:       site.Icon,
			Title:       site.Title,
			Url:         site.URL,
			Category:    categoryName,
			CategoryId:  site.CategoryID,
			Description: site.Description,
			IsUsed:      site.IsUsed,
			CreatedAt:   site.CreatedAt.Format(time.DateTime),
			UpdatedAt:   site.UpdatedAt.Format(time.DateTime),
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
