/**
 * @Author: chentong
 * @Date: 2025/01/31 12:12
 */

package gormx

import (
	"strings"

	"gorm.io/gen/field"
)

func FieldIsDesc(field string) (bool, string) {
	if strings.HasPrefix(field, "-") {
		return true, field[1:]
	}
	return false, field
}

func LikeInner(s string) string {
	return "%" + s + "%"
}

func LikeLeft(s string) string {
	return "%" + s
}

func LikeRight(s string) string {
	return s + "%"
}

// ColumnName 获取 model 列名
func ColumnName(field field.Field) string {
	return field.ColumnName().String()
}
