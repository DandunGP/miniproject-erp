package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Barang_Masuk struct {
	gorm.Model
	No_BPB              string `json:"no_bpb" form:"no_bpb"`
	Tgl_terima          string `json:"tgl_terima" form:"tgl_terima"`
	Qty                 int    `json:"qty" form:"qty"`
	Keterangan          string `json:"keterangan" form:"keterangan"`
	GudangID            int
	Gudang              Gudang `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Persediaan_BarangID int
	Persediaan_Barang   Persediaan_Barang `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
