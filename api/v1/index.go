/**
 * @Author: chentong
 * @Date: 2024/05/26 上午1:50
 */

package v1

import "github.com/ch3nnn/webstack-go/internal/dal/model"

type TreeNode struct {
	Id    int         // 节点ID
	Pid   int         // 父节点ID
	Name  string      // 节点名称
	Icon  string      // 图标
	Sort  int         // 排序
	Child []*TreeNode // 获取子节点切片
}

type CategorySite struct {
	Category string         // 分类
	SiteList []model.StSite // 站点列表
}

type About struct {
	AboutSite   string `json:"about_site"`   // 关于站点
	AboutAuthor string `json:"about_author"` // 关于作者
	IsAbout     bool   `json:"is_about"`     // 是否开启关于
}

type ConfigSite struct {
	SiteTitle   string `json:"site_title"`   // 站点标题
	SiteKeyword string `json:"site_keyword"` // 站点关键字
	SiteDesc    string `json:"site_desc"`    // 站点描述
	SiteRecord  string `json:"site_record"`  // 站点备案
	SiteURL     string `json:"site_url"`     // 备案url
	SiteLogo    string `json:"site_logo"`    // 站点logo
	SiteFavicon string `json:"site_favicon"` // 站点favicon
}

type IndexResp struct {
	About         *About          // 关于页面
	ConfigSite    *ConfigSite     // 站点配置
	CategoryTree  []*TreeNode     // 分类树
	CategorySites []*CategorySite // 归类站点数据
}

type AboutResp struct {
	About
}
