package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Gudang struct {
	gorm.Model
	Kode_gudang string `json:"kode_gudang" form:"kode_gudang"`
	Nama        string `json:"nama" form:"nama"`
	Keterangan  string `json:"keterangan" form:"keterangan"`
}
