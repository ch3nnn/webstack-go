package category

import "time"

// Category 站分类
//go:generate gormgen -structs Category -input .
type Category struct {
	Id         int32  //
	ParentId   int32  //
	Sort       int32  //
	Title      string //
	Icon       string //
	Level      int32
	IsUsed     int32     // 是否启用 1:是 -1:否
	CreateTime time.Time `gorm:"autoCreateTime:true"`
	UpdateTime time.Time `gorm:"autoUpdateTime:true"` //
}
