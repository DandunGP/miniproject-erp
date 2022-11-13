package controllers

import (
	"encoding/json"
	"erp/config"
	"erp/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func GetWOsController(c echo.Context) error {
	var wo []models.Work_order

	if err := config.DB.Preload("BOM").Preload("Gudang").Preload("Officer").Find(&wo).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all wo",
		"wo":     wo,
	})
}

func GetWOController(c echo.Context) error {
	var wo models.Work_order

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("BOM").Preload("Gudang").Preload("Officer").First(&wo).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get wo",
		"wo":      wo,
	})
}

func CreateWOController(c echo.Context) error {
	var wo models.Work_order
	var bom models.Bill_material
	var persediaan models.Invetory

	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	tgl_wo := input["tgl_wo"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, tgl_wo)
	input["tgl_wo"] = tgl
	input["created_at"] = time.Now()
	input["updated_at"] = time.Now()

	if err := config.DB.Model(&bom).Where("id = ?", input["bom_id"]).First(&bom).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&wo).Create(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", bom.Persediaan_BarangID).Updates(map[string]interface{}{
		"qty": gorm.Expr("qty - ?", input["qty_wo"])}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new wo",
		"wo":      input,
	})
}

func UpdateWOController(c echo.Context) error {
	var wo models.Work_order

	id, _ := strconv.Atoi(c.Param("id"))

	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	tgl_wo := input["tgl_wo"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, tgl_wo)
	input["tgl_wo"] = tgl
	input["updated_at"] = time.Now()

	if err := config.DB.Model(&wo).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteWOController(c echo.Context) error {
	var wo models.Work_order

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&wo, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete wo",
	})
}
