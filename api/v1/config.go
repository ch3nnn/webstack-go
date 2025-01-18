/**
 * @Author: chentong
 * @Date: 2025/01/17 19:36
 */

package v1

type ConfigResp struct {
	ID          int    `json:"id"`
	AboutSite   string `json:"about_site"`
	AboutAuthor string `json:"about_author"`
	IsAbout     bool   `json:"is_about"`
	SiteTitle   string `json:"site_title"`
	SiteKeyword string `json:"site_keyword"`
	SiteDesc    string `json:"site_desc"`
	SiteRecord  string `json:"site_record"`
}

type (
	ConfigUpdateReq struct {
		AboutSite   *string `json:"about_site" form:"about_site" `
		AboutAuthor *string `json:"about_author" form:"about_author"`
		IsAbout     *bool   `json:"is_about" form:"is_about"`
		SiteTitle   *string `json:"site_title" form:"site_title"`
		SiteKeyword *string `json:"site_keyword" form:"site_keyword"`
		SiteDesc    *string `json:"site_desc" form:"site_desc"`
		SiteRecord  *string `json:"site_record" form:"site_record"`
	}

	ConfigUpdateResp struct {
		ID int `json:"id"` // 主键ID
	}
)
