/**
 * @Author: chentong
 * @Date: 2025/01/18 14:21
 */

package config

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func (s *service) Update(ctx *gin.Context, req *v1.ConfigUpdateReq) (resp *v1.ConfigUpdateResp, err error) {
	update := make(map[string]any)
	if req.SiteTitle != nil {
		update["site_title"] = *req.SiteTitle
	}
	if req.SiteDesc != nil {
		update["site_desc"] = *req.SiteDesc
	}
	if req.SiteKeyword != nil {
		update["site_keyword"] = *req.SiteKeyword
	}
	if req.SiteRecord != nil {
		update["site_record"] = *req.SiteRecord
	}
	if req.AboutSite != nil {
		update["about_site"] = *req.AboutSite
	}
	if req.AboutAuthor != nil {
		update["about_author"] = *req.AboutAuthor
	}
	if req.IsAbout != nil {
		update["is_about"] = *req.IsAbout
	}

	if _, err = s.configRepo.WithContext(ctx).Update(update, s.configRepo.WhereByID(1)); err != nil {
		return nil, err
	}

	return nil, nil
}
