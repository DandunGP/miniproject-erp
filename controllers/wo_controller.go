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

func GetWOsController(c echo.Context) error {
	var wo []models.WO

	if err := config.DB.Find(&wo).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all wo",
		"wo":     wo,
	})
}

func GetWOController(c echo.Context) error {
	var wo models.WO

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&wo).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get wo",
		"wo":      wo,
	})
}

func CreateWOController(c echo.Context) error {
	var wo models.WO
	var bom models.BOM
	var persediaan models.Persediaan_Barang

	qtyInt, _ := strconv.Atoi(c.FormValue("qty_wo"))
	bomInt, _ := strconv.Atoi(c.FormValue("bom_id"))
	gudangInt, _ := strconv.Atoi(c.FormValue("gudang_id"))
	pegawaiInt, _ := strconv.Atoi(c.FormValue("pegawai_id"))

	wo.NoWO = c.FormValue("no_wo")
	wo.Keterangan = c.FormValue("keterangan")
	wo.TipeJadi = c.FormValue("tipe")
	wo.Qty_WO = qtyInt
	wo.BOMID = bomInt
	wo.GudangID = gudangInt
	wo.PegawaiID = pegawaiInt

	tgl_terima := c.FormValue("tgl_wo")
	dateFormat := "02/01/2006 MST"
	value := tgl_terima + " WIB"
	tgl, _ := time.Parse(dateFormat, value)
	wo.Tgl_WO = tgl

	if err := config.DB.Model(&bom).Where("id = ?", bomInt).First(&bom).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Create(&wo).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", bom.Persediaan_BarangID).Updates(map[string]interface{}{
		"qty": gorm.Expr("qty - ?", qtyInt)}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new wo",
		"wo":      wo,
	})
}

func UpdateWOController(c echo.Context) error {
	var wo models.WO

	qtyInt, _ := strconv.Atoi(c.FormValue("qty_wo"))
	bomInt, _ := strconv.Atoi(c.FormValue("bom_id"))
	gudangInt, _ := strconv.Atoi(c.FormValue("gudang_id"))
	pegawaiInt, _ := strconv.Atoi(c.FormValue("pegawai_id"))

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.WO

	input.NoWO = c.FormValue("no_wo")
	input.Keterangan = c.FormValue("keterangan")
	input.TipeJadi = c.FormValue("tipe")
	input.Qty_WO = qtyInt
	input.BOMID = bomInt
	input.GudangID = gudangInt
	input.PegawaiID = pegawaiInt

	tgl_terima := c.FormValue("tgl_wo")
	dateFormat := "02/01/2006 MST"
	value := tgl_terima + " WIB"
	tgl, _ := time.Parse(dateFormat, value)
	input.Tgl_WO = tgl

	if err := config.DB.Model(&wo).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteWOController(c echo.Context) error {
	var wo models.WO

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&wo, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete wo",
	})
}
