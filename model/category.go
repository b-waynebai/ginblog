package model

import "gorm.io/gorm"

type  Category struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(20);not null" json:"name"`
}