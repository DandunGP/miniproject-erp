package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Invetory struct {
	gorm.Model
	Kode_barang    string `json:"kode_barang" form:"kode_barang"`
	Nama           string `json:"nama" form:"nama"`
	Tipe           string `json:"tipe" form:"tipe"`
	Kategori       string `json:"kategori" form:"kategori"`
	Merek          string `json:"merek" form:"merek"`
	Part_number    string `json:"part" form:"part"`
	Status         string `json:"status" form:"status"`
	Qty            int    `json:"qty" form:"qty"`
	Last_input_qty int    `json:"last" form:"last"`
	Last_input_id  int    `json:"last_id" form:"last_id"`
}
