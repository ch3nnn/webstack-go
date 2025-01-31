/**
 * @Author: chentong
 * @Date: 2024/05/27 上午10:24
 */

package v1

import "time"

type Category struct {
	ID        int        `json:"id"`         // 主键ID
	ParentID  int        `json:"parent_id"`  // 父级分类ID
	Sort      int        `json:"sort"`       // 排序
	Title     string     `json:"title"`      // 名称
	Icon      string     `json:"icon"`       // 图标
	CreatedAt *time.Time `json:"created_at"` // 创建时间
	UpdatedAt *time.Time `json:"updated_at"` // 更新时间
	IsUsed    bool       `json:"is_used"`    // 是否启用 1:是 0:否
	Level     int32      `json:"level"`      // 分类等级
}

type (
	CategoryCreateReq struct {
		ParentID int    `form:"parent_id"` // 分类父ID
		Level    int32  `form:"level"`     // 分类等级 1 一级分类  2 二级分类
		Name     string `form:"name"`      // 菜单名称
		Icon     string `form:"icon"`      // 图标
		IsUsed   bool   `form:"is_used"`   // 是否启用 1:是 0:否
		SortID   int    `form:"sort_id"`   // 排序 ID
	}

	CategoryCreateResp struct {
		Category // 分类信息
	}
)

type (
	CategoryList struct {
		Id     int    `json:"id"`      // ID
		Pid    int    `json:"pid"`     // 父类ID
		Name   string `json:"name"`    // 菜单名称
		Link   string `json:"link"`    // 链接地址
		Icon   string `json:"icon"`    // 图标
		IsUsed bool   `json:"is_used"` // 是否启用 1=启用 0=禁用
		Sort   int    `json:"sort"`    // 排序
		Level  int32  `json:"level"`   // 分类等级 1 一级分类  2 二级分类
	}

	CategoryListReq struct{}

	CategoryListResp struct {
		List []CategoryList `json:"list"` // 分类列表
	}
)

type (
	CategoryDeleteReq struct {
		ID int `uri:"id" binding:"required"` // ID
	}

	CategoryDeleteResp struct{}
)

type (
	CategoryDetailReq struct {
		ID int `uri:"id" binding:"required"` // ID
	}

	CategoryDetailResp struct {
		Id     int    `json:"id"`      // 主键ID
		Pid    int    `json:"pid"`     // 父类ID
		Name   string `json:"name"`    // 分类名称
		Icon   string `json:"icon"`    // 图标
		IsAdd  bool   `json:"is_add"`  // 是否新增子分类
		SortID int    `json:"sort_id"` // 排序 ID
	}
)

type (
	CategoryUpdateReq struct {
		ID     int     `form:"id" binding:"required"` // ID
		Pid    *int    `form:"parent_id"`             // 父类ID
		Name   *string `form:"name"`                  // 菜单名称
		Icon   *string `form:"icon"`                  // 图标
		IsUsed *bool   `form:"used"`                  // 是否启用
		Sort   *int    `form:"sort"`                  // 排序
	}

	CategoryUpdateResp struct {
		Category // 分类信息
	}
)
