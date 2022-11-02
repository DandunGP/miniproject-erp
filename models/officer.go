package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Nama          string    `json:"nama" form:"nama"`
	NIK           string    `json:"nik" form:"nik"`
	Tgl_lahir     time.Time `json:"tgl_lahir" form:"tgl_lahir"`
	Jenis_kelamin string    `json:"jenis_kelamin" form:"jenis_kelamin"`
	Alamat        string    `json:"alamat" form:"alamat"`
	No_hp         string    `json:"no_hp" form:"no_hp"`
	Jabatan       string    `json:"jabatan" form:"jabatan"`
	UserID        int       `json:"user_id" form:"user_id"`
	User          User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
