/**
 * @Author: chentong
 * @Date: 2025/01/18 14:21
 */

package config

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

const (
	LogoWidth     = 200
	LogoHeight    = 50
	FaviconWidth  = 64
	FaviconHeight = 64
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
	if req.LogFile != nil {
		base64Str, err := tools.ResizeMultipartImgToBase64(req.LogFile, LogoWidth, LogoHeight)
		if err != nil {
			base64Str = repository.DefaultLogoBase64
		}

		update["site_logo"] = base64Str
	}
	if req.FaviconFile != nil {
		base64Str, err := tools.ResizeMultipartImgToBase64(req.FaviconFile, FaviconWidth, FaviconHeight)
		if err != nil {
			base64Str = repository.DefaultFaviconBase64
		}

		update["site_favicon"] = base64Str
	}

	if _, err = s.configRepo.WithContext(ctx).Update(update, s.configRepo.WhereByID(1)); err != nil {
		return nil, err
	}

	return nil, nil
}
