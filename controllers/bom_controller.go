package controllers

import (
	"erp/config"
	"erp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBOMsController(c echo.Context) error {
	var boms []models.BOM

	if err := config.DB.Find(&boms).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all boms",
		"boms":   boms,
	})
}

func GetBOMController(c echo.Context) error {
	var bom models.BOM

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&bom).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get bom",
		"bom":     bom,
	})
}

func CreateBOMController(c echo.Context) error {
	var bom models.BOM
	c.Bind(&bom)

	if err := config.DB.Create(&bom).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new bom",
		"bom":     bom,
	})
}

func UpdateBOMController(c echo.Context) error {
	var bom models.BOM

	id, _ := strconv.Atoi(c.Param("id"))

	var input models.BOM
	c.Bind(&input)

	if err := config.DB.Model(&bom).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteBOMController(c echo.Context) error {
	var bom models.BOM

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&bom, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete BOM",
	})
}
