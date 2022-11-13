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

func GetBarangsController(c echo.Context) error {
	var barangs []models.Incoming_goods

	if err := config.DB.Preload("Gudang").Preload("Persediaan_Barang").Find(&barangs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get all barangs",
		"Barangs": barangs,
	})
}

func GetBarangController(c echo.Context) error {
	var barang models.Incoming_goods

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
	var barang models.Incoming_goods
	var persediaan models.Invetory
	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	tgl_terima := input["tgl_terima"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, tgl_terima)
	input["tgl_terima"] = tgl
	input["created_at"] = time.Now()
	input["updated_at"] = time.Now()

	if err := config.DB.Model(&barang).Create(&input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", input["persediaan_barang_id"]).Updates(map[string]interface{}{
		"qty":            gorm.Expr("qty + ?", input["qty"]),
		"last_input_qty": input["qty"],
		"last_input_id":  input["persediaan_barang_id"]}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new Barang",
		"Barang":  input,
	})
}

func UpdateBarangController(c echo.Context) error {
	var barang models.Incoming_goods
	var persediaan models.Invetory
	var input map[string]interface{}

	id, _ := strconv.Atoi("id")

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	tgl_terima := input["tgl_terima"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, tgl_terima)
	input["tgl_terima"] = tgl
	input["updated_at"] = time.Now()

	if err := config.DB.Model(&barang).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", input["persediaan_barang_id"]).Where("last_input_id = ?", input["persediaan_barang_id"]).Updates(map[string]interface{}{"qty": gorm.Expr("qty + ? - last_input_qty", input["qty"])}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}

	if err := config.DB.Model(&persediaan).Where("id = ?", input["persediaan_barang_id"]).Where("last_input_id = ?", input["persediaan_barang_id"]).Updates(map[string]interface{}{"last_input_qty": input["qty"], "last_input_id": input["persediaan_barang_id"]}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	}

	// if err := config.DB.Model(&persediaan).Where("id = ?", input["persediaan_barang_id"]).Where("last_input_id != ?", input["persediaan_barang_id"]).Updates(map[string]interface{}{"qty": gorm.Expr("qty - last_input_qty")}).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	// }

	// if err := config.DB.Model(&persediaan).Where("id = ?", input["persediaan_barang_id"]).Where("last_input_id != ?", input["persediaan_barang_id"]).Updates(map[string]interface{}{"last_input_qty": input["qty"], "last_input_id": input["persediaan_barang_id"]}).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Qty error")
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteBarangController(c echo.Context) error {
	var barang models.Incoming_goods

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&barang, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete barang",
	})
}
