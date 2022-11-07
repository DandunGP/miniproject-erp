package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	UserID        int       `json:"user_id"`
	User          User
}

// func (officer *Officer) BeforeCreate(scope *gorm.Scope) (err error) {
// 	dateFormat := "02/01/2006 MST"
// 	value := officer.Tgl_lahir.String() + " WIB"
// 	tgl, _ := time.Parse(dateFormat, value)
// 	scope.SetColumn("Tgl_lahir", tgl)
// 	return
// }
