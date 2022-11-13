package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Kode_produk string `json:"kode_produk"`
	Nama_produk string `json:"nama_produk"`
	Qty         int    `json:"qty"`
	GudangID    int    `json:"gudang_id"`
	Gudang      Storage
}
