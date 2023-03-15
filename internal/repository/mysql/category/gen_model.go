package category

import "time"

// Category 站分类
//go:generate gormgen -structs Category -input .
type Category struct {
	Id        int32     //
	ParentId  int32     //
	Sort      int32     // 排序
	Title     string    // 名称
	Icon      string    // 图标
	Level     int32     // 分类等级
	IsUsed    int32     // 是否启用 1:是 -1:否
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
