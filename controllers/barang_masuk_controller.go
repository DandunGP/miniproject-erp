package controllers

import (
	"erp/config"
	"erp/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetBarangsController(c echo.Context) error {
	var barangs []models.Barang_Masuk

	if err := config.DB.Preload("Gudang").Preload("Persediaan_Barang").Find(&barangs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get all barangs",
		"Barangs": barangs,
	})
}

func GetBarangController(c echo.Context) error {
	var barang models.Barang_Masuk

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("Gudang").Preload("Persediaan_Barang").First(&barang).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get barang",
		"Barang":  barang,
	})
}

func CreateBarangController(c echo.Context) error {
	var barang models.Barang_Masuk
	var persediaan models.Persediaan_Barang

	qtyInt, _ := strconv.Atoi(c.FormValue("qty"))
	gudang_idInt, _ := strconv.Atoi(c.FormValue("gudang_id"))
	persediaan_barang_idInt, _ := strconv.Atoi(c.FormValue("persediaan_barang_id"))

	barang.No_BPB = c.FormValue("no_bpb")
	barang.Qty = qtyInt
	barang.Keterangan = c.FormValue("keterangan")
	barang.GudangID = gudang_idInt
	barang.Persediaan_BarangID = persediaan_barang_idInt

	tgl_terima := c.FormValue("tgl_terima")
	dateFormat := "02/01/2006 MST"
	value := tgl_terima + " WIB"
	tgl, _ := time.Parse(dateFormat, value)
	barang.Tgl_terima = tgl

	if err := config.DB.Create(&barang).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", persediaan_barang_idInt).Updates(map[string]interface{}{
		"qty":            gorm.Expr("qty + ?", qtyInt),
		"last_input_qty": qtyInt,
		"last_input_id":  persediaan_barang_idInt}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Barang",
		"Barang":  barang,
	})
}

func UpdateBarangController(c echo.Context) error {
	var barang models.Barang_Masuk
	var persediaan models.Persediaan_Barang

	id, _ := strconv.Atoi(c.Param("id"))

	no_bpb := c.FormValue("no_bpb")
	qty := c.FormValue("qty")
	keterangan := c.FormValue("keterangan")
	gudang_id := c.FormValue("gudang_id")
	persediaan_barang_id := c.FormValue("persediaan_barang_id")
	qtyInt, _ := strconv.Atoi(qty)
	gudang_idInt, _ := strconv.Atoi(gudang_id)
	persediaan_barang_idInt, _ := strconv.Atoi(persediaan_barang_id)

	var input models.Barang_Masuk

	tgl_terima := c.FormValue("tgl_terima")
	dateFormat := "02/01/2006 MST"
	value := tgl_terima + " WIB"
	tgl, _ := time.Parse(dateFormat, value)
	barang.Tgl_terima = tgl

	input.No_BPB = no_bpb
	input.Tgl_terima = tgl
	input.Qty = qtyInt
	input.Keterangan = keterangan
	input.GudangID = gudang_idInt
	input.Persediaan_BarangID = persediaan_barang_idInt

	if err := config.DB.Model(&barang).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", persediaan_barang_idInt).Where("last_input_id = ?", persediaan_barang_idInt).Updates(map[string]interface{}{"qty": gorm.Expr("qty + ? - last_input_qty", qtyInt)}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", persediaan_barang_idInt).Where("last_input_id = ?", persediaan_barang_idInt).Updates(map[string]interface{}{"last_input_qty": qtyInt, "last_input_id": persediaan_barang_idInt}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}

	// if err := config.DB.Model(&persediaan).Where("id = ?", persediaan_barang_idInt).Where("last_input_id != ?", persediaan_barang_idInt).Updates(map[string]interface{}{"qty": gorm.Expr("qty - last_input_qty")}).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	// }

	// if err := config.DB.Model(&persediaan).Where("id = ?", persediaan_barang_idInt).Where("last_input_id != ?", persediaan_barang_idInt).Updates(map[string]interface{}{"last_input_qty": qtyInt, "last_input_id": persediaan_barang_idInt}).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteBarangController(c echo.Context) error {
	var barang models.Barang_Masuk

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&barang, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete barang",
	})
}
