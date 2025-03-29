/**
 * @Author: chentong
 * @Date: 2025/01/17 19:36
 */

package v1

import "mime/multipart"

type ConfigResp struct {
	ID          int    `json:"id"`           // 主键ID
	AboutSite   string `json:"about_site"`   // 关于网站的信息
	AboutAuthor string `json:"about_author"` // 关于作者的信息
	IsAbout     bool   `json:"is_about"`     // 是否显示关于信息
	SiteTitle   string `json:"site_title"`   // 网站标题
	SiteKeyword string `json:"site_keyword"` // 网站关键词
	SiteDesc    string `json:"site_desc"`    // 网站描述
	SiteRecord  string `json:"site_record"`  // 网站备案号
	SiteURL     string `json:"site_url"`     // 网站备案管理 url
	SiteLogo    string `json:"site_logo"`    // 网站Logo URL
	SiteFavicon string `json:"site_favicon"` // 网站Favicon URL
}

type (
	ConfigUpdateReq struct {
		AboutSite   *string               `json:"about_site" form:"about_site"`     // 关于网站的信息
		AboutAuthor *string               `json:"about_author" form:"about_author"` // 关于作者的信息
		IsAbout     *bool                 `json:"is_about" form:"is_about"`         // 是否显示关于信息
		SiteTitle   *string               `json:"site_title" form:"site_title"`     // 网站标题
		SiteKeyword *string               `json:"site_keyword" form:"site_keyword"` // 网站关键词
		SiteDesc    *string               `json:"site_desc" form:"site_desc"`       // 网站描述
		SiteRecord  *string               `json:"site_record" form:"site_record"`   // 网站备案号
		SiteURL     *string               `json:"site_url" form:"site_url"`         // 网站备案管理 url
		LogoFile    *multipart.FileHeader `json:"logo" form:"logo"`                 // 上传 logo 图片
		FaviconFile *multipart.FileHeader `json:"favicon" form:"favicon"`           // 上传 favicon 图片
	}

	ConfigUpdateResp struct {
		ID int `json:"id"` // 主键ID
	}
)
