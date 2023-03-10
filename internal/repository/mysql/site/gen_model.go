package site

import (
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/category"
	"time"
)

// Site 网站信息
//go:generate gormgen -structs Site -input .
type Site struct {
	Id          int32 //
	CategoryId  int32 //
	Category    category.Category
	IsUsed      int32     // 是否启用 1:是  -1:否
	Title       string    //
	Thumb       string    //
	Description string    //
	Url         string    //
	CreateTime  time.Time `gorm:"autoCreateTime:true"`
	UpdateTime  time.Time `gorm:"autoUpdateTime:true"` //
}
