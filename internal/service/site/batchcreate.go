/**
 * @Author: chentong
 * @Date: 2024/06/04 下午4:33
 */

package site

import (
	"context"
	"strings"

	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/validator"
	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

func (s *service) parseURL(u string) (urls []string) {
	for _, u := range strings.Split(u, "\n") {
		if validator.IsUrl(u) || validator.IsIp(u) || validator.IsIpPort(u) {
			urls = append(urls, u)
		}
	}

	return
}

func (s *service) BatchCreate(ctx context.Context, req *v1.SiteCreateReq) (*v1.SiteCreateResp, error) {
	workerPool := tools.NewWorkerPool(5, 20)
	workerPool.Start()

	var successCnt int
	var failURLs []string
	for _, u := range s.parseURL(req.Url) {
		workerPool.AddJob(func() {
			var (
				g                 errgroup.Group
				title, icon, desc string
			)

			u = strings.TrimSpace(u)

			g.Go(func() (err error) {
				title, err = getWebTitle(u)
				return
			})
			g.Go(func() (err error) {
				icon, err = getWebLogoIconBase64(u)
				return
			})
			g.Go(func() (err error) {
				desc, err = getWebDescription(u)
				return
			})

			if err := g.Wait(); err != nil {
				if !req.FailSwitch {
					failURLs = append(failURLs, u)
					return
				}
			}

			_, err := s.siteRepository.WithContext(ctx).Create(&model.StSite{
				Title:       condition.Ternary(title != "", title, u),
				Icon:        icon,
				Description: desc,
				URL:         u,
				CategoryID:  req.CategoryID,
				IsUsed:      req.IsUsed,
				Sort:        0,
			})
			if err != nil {
				failURLs = append(failURLs, u)
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
