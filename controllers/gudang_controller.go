package controllers

import (
	"erp/config"
	"erp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetGudangsController(c echo.Context) error {
	var gudangs []models.Gudang

	if err := config.DB.Find(&gudangs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get all gudangs",
		"gudangs": gudangs,
	})
}

func GetGudangController(c echo.Context) error {
	var gudang models.Gudang

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&gudang).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get gudang",
		"Gudang":  gudang,
	})
}

func CreateGudangController(c echo.Context) error {
	var gudang models.Gudang
	c.Bind(&gudang)

	if err := config.DB.Create(&gudang).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new gudang",
		"gudang":  gudang,
	})
}

func UpdateGudangController(c echo.Context) error {
	var gudang models.Gudang

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Gudang
	c.Bind(&input)

	if err := config.DB.Model(&gudang).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteGudangController(c echo.Context) error {
	var gudang models.Gudang

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&gudang, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete gudang",
	})
}
