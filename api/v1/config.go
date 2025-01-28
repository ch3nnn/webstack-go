/**
 * @Author: chentong
 * @Date: 2025/01/17 19:36
 */

package v1

import "mime/multipart"

type ConfigResp struct {
	ID          int    `json:"id"`
	AboutSite   string `json:"about_site"`
	AboutAuthor string `json:"about_author"`
	IsAbout     bool   `json:"is_about"`
	SiteTitle   string `json:"site_title"`
	SiteKeyword string `json:"site_keyword"`
	SiteDesc    string `json:"site_desc"`
	SiteRecord  string `json:"site_record"`
	SiteLogo    string `json:"site_logo"`
	SiteFavicon string `json:"site_favicon"`
}

type (
	ConfigUpdateReq struct {
		AboutSite   *string               `json:"about_site" form:"about_site" `
		AboutAuthor *string               `json:"about_author" form:"about_author"`
		IsAbout     *bool                 `json:"is_about" form:"is_about"`
		SiteTitle   *string               `json:"site_title" form:"site_title"`
		SiteKeyword *string               `json:"site_keyword" form:"site_keyword"`
		SiteDesc    *string               `json:"site_desc" form:"site_desc"`
		SiteRecord  *string               `json:"site_record" form:"site_record"`
		LogFile     *multipart.FileHeader `json:"log"`     // 上传 logo 图片
		FaviconFile *multipart.FileHeader `json:"favicon"` // 上传 favicon 图片

	}

	ConfigUpdateResp struct {
		ID int `json:"id"` // 主键ID
	}
)
