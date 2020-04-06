package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Mobile   string
	Password string
}

// 表名
func (User) TableName() string {
	return "user"
}
