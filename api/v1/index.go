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

type IndexResponseData struct {
	CategoryTree  []*TreeNode     // 分类树
	CategorySites []*CategorySite // 归类站点数据
}
