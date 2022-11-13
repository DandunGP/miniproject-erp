package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Incoming_goods struct {
	gorm.Model
	No_BPB              string    `json:"no_bpb" form:"no_bpb"`
	Tgl_terima          time.Time `json:"tgl_terima" form:"tgl_terima"`
	Qty                 int       `json:"qty" form:"qty"`
	Keterangan          string    `json:"keterangan" form:"keterangan"`
	GudangID            int       `json:"gudang_id" form:"gudang_id"`
	Gudang              Storage
	Persediaan_BarangID int `json:"persediaan_barang_id" form:"persediaan_barang_id"`
	Persediaan_Barang   Invetory
}
