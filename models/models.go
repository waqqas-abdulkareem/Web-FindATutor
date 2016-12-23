package model

import(
	"github.com/jinzhu/gorm"
)

type Subject struct{
	gorm.Model
	Name string	`gorm:"not null; unique; index; size:255"`
}