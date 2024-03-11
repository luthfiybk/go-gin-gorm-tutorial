package model

type Products struct {
	Id int `gorm:"type:int;primary_key"`
	Product_Name string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
}