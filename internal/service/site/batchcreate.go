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
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

func (s *service) parseURL(url string) (urls []string) {
	for _, u := range strings.Split(url, "\n") {
		urls = append(urls, u)
	}

	return
}

func (s *service) BatchCreate(ctx context.Context, req *v1.SiteCreateReq) (*v1.SiteCreateResp, error) {
	workerPool := tools.NewWorkerPool(5, 20)
	workerPool.Start()

	var successCnt int
	var failURLs []string
	for _, url := range s.parseURL(req.Url) {
		workerPool.AddJob(func() {
			var (
				g                 errgroup.Group
				title, icon, desc string
			)

			url := strings.TrimSpace(url)

			g.Go(func() (err error) {
				title, err = getWebTitle(url)
				return
			})
			g.Go(func() (err error) {
				icon, err = getWebLogoIconBase64(url)
				return
			})
			g.Go(func() (err error) {
				desc, err = getWebDescription(url)
				return
			})

			if err := g.Wait(); err != nil {
				failURLs = append(failURLs, url)
				return
			}

			_, err := s.siteRepository.WithContext(ctx).Create(&model.StSite{
				Title:       title,
				Icon:        icon,
				Description: desc,
				URL:         url,
				CategoryID:  req.CategoryID,
				IsUsed:      req.IsUsed,
				Sort:        0,
			})
			if err != nil {
				failURLs = append(failURLs, url)
				return
			}

			successCnt++
		})
	}

	workerPool.Wait()

	return &v1.SiteCreateResp{
		FailCount:    len(failURLs),
		SuccessCount: successCnt,
		FailURLs:     failURLs,
	}, nil
}
