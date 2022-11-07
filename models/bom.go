package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type BOM struct {
	gorm.Model
	Kode_BOM            string `json:"kode_bom" form:"kode_bom"`
	Nama_BOM            string `json:"nama_bom" form:"nama_bom"`
	Keterangan          string `json:"keterangan" form:"keterangan"`
	Persediaan_BarangID int    `json:"persediaan_barang_id"`
	Persediaan_Barang   Persediaan_Barang
	GudangID            int `json:"gudang_id"`
	Gudang              Gudang
}
