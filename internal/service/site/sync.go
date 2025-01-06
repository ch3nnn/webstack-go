/**
 * @Author: chentong
 * @Date: 2024/11/12 16:40
 */

package site

import (
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func (s *service) Sync(ctx *gin.Context, req *v1.SiteSyncReq) (resp *v1.SiteSyncResp, err error) {
	var (
		g                 errgroup.Group
		title, icon, desc string
	)

	site, err := s.siteRepository.WithContext(ctx).FindOne(s.siteRepository.WhereByID(req.ID))
	if err != nil {
		return nil, err
	}

	url := strings.TrimSpace(site.URL)

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
		return nil, err
	}

	_, err = s.siteRepository.WithContext(ctx).Update(&model.StSite{
		Title:       title,
		Icon:        icon,
		Description: desc,
		IsUsed:      false,
	},
		s.siteRepository.WhereByID(req.ID),
	)
	if err != nil {
		return nil, err
	}

	return &v1.SiteSyncResp{ID: site.ID}, nil
}
