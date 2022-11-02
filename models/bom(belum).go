package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type BOM struct {
	gorm.Model
	Kode_BOM   string
	Nama_BOM   string
	Keterangan string
	Barang     string
	Title      string `json:"title" form:"Title"`
	Publisher  string `json:"publisher" form:"Publisher"`
	Year       int    `json:"year" form:"Year"`
}
