package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/query"
	"mime/multipart"
	"path"
)

type UpdateSiteRequest struct {
	Id          int64                 `json:"id"`
	CategoryId  int64                 `json:"category_id"` // 网站分类id
	Title       string                `json:"title"`       // 网站标题
	Thumb       string                `json:"thumb"`       // 网站 logo
	Description string                `json:"description"` // 网站描述
	Url         string                `json:"url"`         // 网站地址
	File        *multipart.FileHeader `json:"file"`        // 上传 logo 图片
}

func getWebThumb(ctx core.Context, updateSite *UpdateSiteRequest) (thumb string) {
	// 修改操作是否手动上传 logo 图片
	if updateSite.File != nil {
		return getWebLogoIconUrlByUploadImg(ctx)
	}
	return updateSite.Thumb
}

func getWebLogoIconUrlByUploadImg(ctx core.Context) string {
	file, _ := ctx.FormFile("file")
	dst := path.Join("/upload/" + file.Filename) // 上传静态文件 url
	if err := ctx.SaveUploadedFile(file, path.Join("assets", dst)); err != nil {
		return ""
	}
	return dst
}

func (s *service) UpdateSite(ctx core.Context, updateSite *UpdateSiteRequest) (err error) {
	if _, err = query.Site.WithContext(ctx.RequestContext()).
		Where(query.Site.ID.Eq(updateSite.Id)).
		Updates(map[string]any{
			"CategoryID":  updateSite.CategoryId,
			"Title":       updateSite.Title,
			"Thumb":       getWebThumb(ctx, updateSite),
			"Description": updateSite.Description,
			"Url":         updateSite.Url,
		}); err != nil {
		return err
	}
	return
}
