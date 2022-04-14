package model

type Category struct {
	Id   int    `gorm:"column:id;type:int(20);not null" json:"id"`
	Name string `gorm:"column:name;type:varchar(20);not null" json:"name"`
}
