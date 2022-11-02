package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"Title"`
	Publisher string `json:"publisher" form:"Publisher"`
	Year      int    `json:"year" form:"Year"`
}
