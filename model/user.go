package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:username;type:varchar(20);not null" json:"username"`
	PassWord string `gorm:"column:password;type:varchar(20);not null" json:"password"`
	Role     int `gorm:"column:role;type:int(20);not null" json:"role"`
}
