package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(20);not null" json:"name"`
	Title string `gorm:"column:title;type:varchar(20);not null" json:"title"`
	Cid int `gorm:"column:cid;type:int(20);not null" json:"cid"`
	Desc string `gorm:"column:name;type:varchar(200);not null" json:"desc"`
	Content string `gorm:"column:Content;type:longtext" json:"content"`
	Img string `gorm:"column:img;type:varchar(20)" json:"img"`
}