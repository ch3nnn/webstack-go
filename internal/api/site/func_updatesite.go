package site

import (
	"github.com/ch3nnn/webstack-go/internal/code"
	"github.com/ch3nnn/webstack-go/internal/pkg/core"
	"github.com/ch3nnn/webstack-go/internal/services/site"
	"mime/multipart"
	"net/http"
)

type updateSiteRequest struct {
	Id          int64                 `form:"id,omitempty"`
	CategoryId  int64                 `form:"category_id,omitempty"` // 网站分类id
	Title       string                `form:"title,omitempty"`       // 网站标题
	Thumb       string                `form:"thumb,omitempty"`       // 网站 logo
	Description string                `form:"description,omitempty"` // 网站描述
	Url         string                `form:"url,omitempty"`         // 网站地址
	File        *multipart.FileHeader `json:"file"`                  // 上传 logo 图片
}

type updateSiteResponse struct {
	Id int64 `json:"id"` // 主键ID
}

// UpdateSite 编辑网站
// @Summary 编辑网站
// @Description 编辑网站
// @Tags API.site
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request formData updateSiteRequest true "请求信息"
// @Success 200 {object} updateSiteResponse
// @Failure 400 {object} code.Failure
// @Router /api/site [put]
func (h *handler) UpdateSite() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateSiteRequest)
		res := new(updateSiteResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		if file, _ := c.FormFile("file"); file != nil {
			req.File = file
		}
		if err := h.siteService.UpdateSite(c, (*site.UpdateSiteRequest)(req)); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SiteUpdateError,
				code.Text(code.SiteUpdateError)).WithError(err),
			)
			return
		}
		res.Id = req.Id
		c.Payload(res)
	}
}
