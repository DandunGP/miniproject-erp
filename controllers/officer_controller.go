package controllers

import (
	"erp/config"
	"erp/lib/database"
	"erp/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
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

	officer.Nama = c.FormValue("nama")
	officer.NIK = c.FormValue("nik")
	officer.Jenis_kelamin = c.FormValue("jenis_kelamin")
	officer.Alamat = c.FormValue("alamat")
	officer.No_hp = c.FormValue("no_hp")
	officer.Jabatan = c.FormValue("jabatan")
	officer.UserID, _ = strconv.Atoi(c.FormValue("user_id"))

	tgl_lahir := c.FormValue("tgl_lahir")
	dateFormat := "02/01/2006 MST"
	value := tgl_lahir + " WIB"
	tgl, _ := time.Parse(dateFormat, value)
	officer.Tgl_lahir = tgl

	if err := config.DB.Create(&officer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new officer",
		"officer": officer,
	})
}

func UpdateOfficerController(c echo.Context) error {
	var officer models.Officer

	id, _ := strconv.Atoi(c.Param("id"))

	nama := c.FormValue("nama")
	nik := c.FormValue("nik")
	tgl_lahir := c.FormValue("tgl_lahir")
	jenis_kelamin := c.FormValue("jenis_kelamin")
	alamat := c.FormValue("alamat")
	no_hp := c.FormValue("no_hp")
	jabatan := c.FormValue("jabatan")
	user_id, _ := strconv.Atoi(c.FormValue("user_id"))

	dateFormat := "02/01/2006 MST"
	value := tgl_lahir + " WIB"
	tgl, _ := time.Parse(dateFormat, value)

	var input models.Officer

	input.Nama = nama
	input.NIK = nik
	input.Tgl_lahir = tgl
	input.Jenis_kelamin = jenis_kelamin
	input.Alamat = alamat
	input.No_hp = no_hp
	input.Jabatan = jabatan
	input.UserID = user_id

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
