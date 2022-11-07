package controllers

import (
	"erp/config"
	"erp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetProduksController(c echo.Context) error {
	var produks []models.Produk

	if err := config.DB.Find(&produks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success get all produks",
		"produks": produks,
	})
}

func GetProdukController(c echo.Context) error {
	var produk models.Produk

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&produk).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get produk",
		"produk":  produk,
	})
}

func CreateProdukController(c echo.Context) error {
	var produk models.Produk
	c.Bind(&produk)

	if err := config.DB.Create(&produk).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new produk",
		"produk":  produk,
	})
}

func UpdateProdukController(c echo.Context) error {
	var produk models.Produk

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Produk
	c.Bind(&input)

	if err := config.DB.Model(&produk).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteProdukController(c echo.Context) error {
	var produk models.Produk

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&produk, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete produk",
	})
}
