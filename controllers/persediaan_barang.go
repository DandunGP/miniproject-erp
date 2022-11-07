package controllers

import (
	"erp/config"
	"erp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetPersediaansController(c echo.Context) error {
	var persediaans []models.Persediaan_Barang

	if err := config.DB.Find(&persediaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "success get all persediaan barang",
		"Persediaans": persediaans,
	})
}

func GetPersediaanController(c echo.Context) error {
	var persediaan models.Persediaan_Barang

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&persediaan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get persediaan",
		"Persediaan": persediaan,
	})
}

func CreatePersediaanController(c echo.Context) error {
	var persediaan models.Persediaan_Barang
	c.Bind(&persediaan)

	if err := config.DB.Create(&persediaan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success create new persediaan barang",
		"Persediaan": persediaan,
	})
}

func UpdatePersediaanController(c echo.Context) error {
	var persediaan models.Persediaan_Barang

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.Persediaan_Barang
	c.Bind(&input)

	if err := config.DB.Model(&persediaan).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeletePersediaanController(c echo.Context) error {
	var persediaan models.Persediaan_Barang

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&persediaan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete persediaan",
	})
}
