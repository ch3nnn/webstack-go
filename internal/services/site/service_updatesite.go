package site

import (
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql"
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/site"
	"mime/multipart"
	"path"
)

type UpdateSiteRequest struct {
	Id          int32                 `json:"id"`
	CategoryId  int32                 `json:"category_id"` // 网站分类id
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
	qb := site.NewQueryBuilder()
	qb.WhereId(mysql.EqualPredicate, updateSite.Id)
	data := map[string]any{
		"CategoryId":  updateSite.CategoryId,
		"Title":       updateSite.Title,
		"Thumb":       getWebThumb(ctx, updateSite),
		"Description": updateSite.Description,
		"Url":         updateSite.Url,
	}
	if err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data); err != nil {
		return err
	}
	return
}
