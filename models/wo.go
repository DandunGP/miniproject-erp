package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type WO struct {
	gorm.Model
	NoWO       string    `json:"no_wo" form:"no_wo"`
	Tgl_WO     time.Time `json:"tgl_wo" form:"tgl_wo"`
	Qty_WO     int       `json:"qty_wo" form:"qty_wo"`
	Keterangan string    `json:"keterangan" form:"keterangan"`
	TipeJadi   string    `json:"tipe" form:"tipe"`
	BOMID      int       `json:"bom_id" form:"bom_id"`
	BOM        BOM
	Qty        int `json:"qty" form:"qty"`
	GudangID   int `json:"gudang_id" form:"gudang_id"`
	Gudang     Gudang
}
