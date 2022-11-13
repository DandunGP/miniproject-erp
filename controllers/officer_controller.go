package controllers

import (
	"encoding/json"
	"erp/config"
	"erp/lib/database"
	"erp/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func GetOfficersController(c echo.Context) error {
	officers, err := database.GetOfficers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success get all officers",
		"officers": officers,
	})
}

func GetOfficerController(c echo.Context) error {
	var officer models.Officer

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).Preload("User").First(&officer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get officer",
		"officer": officer,
	})
}

func CreateOfficerController(c echo.Context) error {
	var officer models.Officer
	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	tgl_lahir := input["tgl_lahir"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, tgl_lahir)
	input["tgl_lahir"] = tgl
	input["created_at"] = time.Now()
	input["updated_at"] = time.Now()

	if err := config.DB.Model(&officer).Create(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new officer",
		"officer": input,
	})
}

func UpdateOfficerController(c echo.Context) error {
	var officer models.Officer

	id, _ := strconv.Atoi(c.Param("id"))

	var input map[string]interface{}

	body, _ := ioutil.ReadAll(c.Request().Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	tgl_lahir := input["tgl_lahir"].(string)
	dateFormat := "02/01/2006"
	tgl, _ := time.Parse(dateFormat, tgl_lahir)
	input["tgl_lahir"] = tgl
	input["updated_at"] = time.Now()

	if err := config.DB.Model(&officer).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteOfficerController(c echo.Context) error {
	var officer models.Officer

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&officer, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete officer",
	})
}
