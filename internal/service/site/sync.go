/**
 * @Author: chentong
 * @Date: 2024/11/12 16:40
 */

package site

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/model"
)

func (s *service) Sync(ctx *gin.Context, req *v1.SiteSyncReq) (resp *v1.SiteSyncResp, err error) {
	site, err := s.siteRepository.WithContext(ctx).FindOne(s.siteRepository.WhereByID(req.ID))
	if err != nil {
		return nil, err
	}

	_, err = s.siteRepository.WithContext(ctx).Update(&model.StSite{
		Title:       getWebTitle(site.URL),
		Icon:        getWebLogoIcon(site.URL),
		Description: getWebDescription(site.URL),
		IsUsed:      false,
	},
		s.siteRepository.WhereByID(req.ID),
	)
	if err != nil {
		return nil, err
	}

	return &v1.SiteSyncResp{ID: site.ID}, nil
}
