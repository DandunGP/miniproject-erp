package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Work_order struct {
	gorm.Model
	NoWO       string    `json:"no_wo" form:"no_wo"`
	Tgl_WO     time.Time `json:"tgl_wo" form:"tgl_wo"`
	Qty_WO     int       `json:"qty_wo" form:"qty_wo"`
	Keterangan string    `json:"keterangan" form:"keterangan"`
	TipeJadi   string    `json:"tipe_jadi" form:"tipe"`
	BOMID      int       `json:"bom_id" form:"bom_id"`
	BOM        Bill_material
	GudangID   int `json:"gudang_id" form:"gudang_id"`
	Gudang     Storage
	PegawaiID  int `json:"pegawai_id" form:"pegawai_id"`
}
