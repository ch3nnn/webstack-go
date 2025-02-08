/**
 * @Author: chentong
 * @Date: 2024/06/30 下午10:14
 */

package site

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
	"github.com/ch3nnn/webstack-go/internal/dal/repository"
	"github.com/ch3nnn/webstack-go/pkg/tools"
)

const (
	FaviconWidth  = 64
	FaviconHeight = 64
)

func (s *service) Update(ctx *gin.Context, req *v1.SiteUpdateReq) (resp *v1.SiteUpdateResp, err error) {
	update := make(map[string]any)

	if req.CategoryId != 0 {
		update["CategoryID"] = req.CategoryId
	}
	if req.Title != "" {
		update["Title"] = req.Title
	}
	if req.Icon != "" {
		base64Str, err := tools.ResizeURLImgToBase64(req.Icon, FaviconWidth, FaviconHeight)
		if err != nil {
			base64Str = repository.DefaultFaviconBase64
		}

		update["Icon"] = base64Str
	}
	if req.File != nil {
		base64Str, err := tools.ResizeMultipartImgToBase64(req.File, FaviconWidth, FaviconHeight)
		if err != nil {
			base64Str = repository.DefaultFaviconBase64
		}

		update["Icon"] = base64Str
	}
	if req.Description != "" {
		update["Description"] = req.Description
	}
	if req.Url != "" {
		update["Url"] = req.Url
	}
	if req.IsUsed != nil {
		update["IsUsed"] = req.IsUsed
	}
	if req.Sort >= 0 {
		update["Sort"] = req.Sort
	}

	_, err = s.siteRepository.WithContext(ctx).Update(update, s.siteRepository.WhereByID(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.SiteUpdateResp{ID: req.Id}, nil
}
