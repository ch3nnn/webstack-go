/**
 * @Author: chentong
 * @Date: 2025/01/18 14:21
 */

package config

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/query"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/pkg/gormx"
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
		column := gormx.ColumnName(query.SysConfig.SiteTitle)
		update[column] = *req.SiteTitle
	}
	if req.SiteDesc != nil {
		column := gormx.ColumnName(query.SysConfig.SiteDesc)
		update[column] = *req.SiteDesc
	}
	if req.SiteKeyword != nil {
		column := gormx.ColumnName(query.SysConfig.SiteKeyword)
		update[column] = *req.SiteKeyword
	}
	if req.SiteRecord != nil {
		column := gormx.ColumnName(query.SysConfig.SiteRecord)
		update[column] = *req.SiteRecord
	}
	if req.SiteURL != nil {
		column := gormx.ColumnName(query.SysConfig.SiteURL)
		update[column] = *req.SiteURL
	}
	if req.AboutSite != nil {
		column := gormx.ColumnName(query.SysConfig.AboutSite)
		update[column] = *req.AboutSite
	}
	if req.AboutAuthor != nil {
		column := gormx.ColumnName(query.SysConfig.AboutAuthor)
		update[column] = *req.AboutAuthor
	}
	if req.IsAbout != nil {
		column := gormx.ColumnName(query.SysConfig.IsAbout)
		update[column] = *req.IsAbout
	}
	if req.LogoFile != nil && req.LogoFile.Size > 0 {
		base64Str, err := tools.ResizeMultipartImgToBase64(req.LogoFile, LogoWidth, LogoHeight)
		if err != nil {
			base64Str = repository.DefaultLogoBase64
		}

		column := gormx.ColumnName(query.SysConfig.SiteLogo)
		update[column] = base64Str
	}
	if req.FaviconFile != nil && req.FaviconFile.Size > 0 {
		base64Str, err := tools.ResizeMultipartImgToBase64(req.FaviconFile, FaviconWidth, FaviconHeight)
		if err != nil {
			base64Str = repository.DefaultFaviconBase64
		}

		column := gormx.ColumnName(query.SysConfig.SiteFavicon)
		update[column] = base64Str
	}

	if _, err = s.configRepo.WithContext(ctx).Update(update, s.configRepo.WhereByID(1)); err != nil {
		return nil, err
	}

	return nil, nil
}
