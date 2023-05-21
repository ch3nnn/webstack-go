package site

import (
	"github.com/ch3nnn/webstack-go/internal/repository/mysql/category"
	"time"
)

type IsUsedStatus int32

// Site 网站信息
//go:generate gormgen -structs Site -input .
type Site struct {
	Id          int32             //
	CategoryId  int32             // 网站分类id
	Category    category.Category // 网站分类
	IsUsed      IsUsedStatus      // 是否启用 1:是  -1:否
	Title       string            // 网站标题
	Thumb       string            // 网站 logo
	Description string            // 网站描述
	Url         string            //  网站地址
	CreatedAt   time.Time         `gorm:"time"` // 创建时间
	UpdatedAt   time.Time         `gorm:"time"` // 更新时间
}

const (
	Off  IsUsedStatus = -1
	Open IsUsedStatus = 1
)
