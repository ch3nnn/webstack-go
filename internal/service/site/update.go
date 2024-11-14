/**
 * @Author: chentong
 * @Date: 2024/06/30 下午10:14
 */

package site

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	v1 "github.com/ch3nnn/webstack-go/api/v1"
)

func IconPath(ctx *gin.Context, req *v1.SiteUpdateReq) string {
	// 修改操作是否手动上传 logo 图片
	if req.File != nil {
		return getWebLogoIconUrlByUploadImg(ctx)
	}

	if req.Icon == nil {
		return ""
	}

	return *req.Icon
}

func getWebLogoIconUrlByUploadImg(ctx *gin.Context) string {
	file, _ := ctx.FormFile("file")
	dst := filepath.Join("upload", file.Filename) // 上传静态文件 url
	if err := ctx.SaveUploadedFile(file, filepath.Join("web", dst)); err != nil {
		return ""
	}
	return filepath.Join("/", dst)
}

func (s *service) Update(ctx *gin.Context, req *v1.SiteUpdateReq) (resp *v1.SiteUpdateResp, err error) {
	update := make(map[string]any)

	if req.CategoryId != nil {
		update["CategoryID"] = req.CategoryId
	}
	if req.Title != nil {
		update["Title"] = req.Title
	}
	if req.Icon != nil {
		update["Icon"] = IconPath(ctx, req)
	}
	if req.Description != nil {
		update["Description"] = req.Description
	}
	if req.Url != nil {
		update["Url"] = req.Url
	}
	if req.IsUsed != nil {
		update["IsUsed"] = req.IsUsed
	}

	_, err = s.siteRepository.WithContext(ctx).Update(update, s.siteRepository.WhereByID(req.Id))
	if err != nil {
		return nil, err
	}

	return &v1.SiteUpdateResp{ID: req.Id}, nil
}
