// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSite = "site"

// Site mapped from table <site>
type Site struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CategoryID  int64     `gorm:"column:category_id;comment:分类id" json:"category_id"`
	Title       string    `gorm:"column:title;comment:网站标题" json:"title"`
	Thumb       string    `gorm:"column:thumb;comment:网站logo" json:"thumb"`
	Description string    `gorm:"column:description;comment:网站描述" json:"description"`
	URL         string    `gorm:"column:url;comment:网站地址" json:"url"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
	IsUsed      int64     `gorm:"column:is_used;default:-1;comment:是否使用" json:"is_used"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
}

// TableName Site's table name
func (*Site) TableName() string {
	return TableNameSite
}
