/**
 * @Author: chentong
 * @Date: 2024/06/04 下午4:33
 */

package site

import (
	"context"
	"strings"

	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func (s *service) parseURL(url string) (urls []string) {
	for _, u := range strings.Split(url, "\n") {
		urls = append(urls, u)
	}

	return
}

func (s *service) BatchCreate(ctx context.Context, req *v1.SiteCreateReq) (*v1.SiteCreateResp, error) {
	var failCnt, successCnt int
	for _, url := range s.parseURL(req.Url) {
		var (
			g                 errgroup.Group
			title, icon, desc string
		)

		url = strings.TrimSpace(url)

		g.Go(func() error {
			title = getWebTitle(url)
			return nil
		})
		g.Go(func() error {
			icon = getWebLogoIcon(url)
			return nil
		})
		g.Go(func() error {
			desc = getWebDescription(url)
			return nil
		})

		if err := g.Wait(); err != nil {
			return nil, err
		}

		_, err := s.siteRepository.WithContext(ctx).Create(&model.StSite{
			Title:       title,
			Icon:        icon,
			Description: desc,
			URL:         url,
			CategoryID:  req.CategoryID,
		})
		if err != nil {
			failCnt++
			continue
		}

		successCnt++
	}

	return &v1.SiteCreateResp{
		FailCount:    failCnt,
		SuccessCount: successCnt,
	}, nil
}
