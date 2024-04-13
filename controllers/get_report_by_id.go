package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetReportByID(c echo.Context) error {
	id := c.Param("id")
	conv, err := strconv.Atoi(id)
	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: "Error converting id",
			Success: false,
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	report := models.Report{}
	err = configs.DB.Preload("Images").Preload("Status").Find(&report, conv).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Paket tidak ditemukan
			BaseResponseError := models.ErrorResponse{
				Message: "Report not found",
				Success: false,
				Code:    http.StatusNotFound,
			}
			return c.JSON(http.StatusNotFound, BaseResponseError)
		}
		// Terjadi kesalahan lain dalam database
		BaseResponseError := models.ErrorResponse{
			Message: "Failed to get report",
			Success: false,
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, BaseResponseError)
	}

	if report.ID == 0 {
		// Paket tidak ditemukan
		BaseResponseError := models.ErrorResponse{
			Message: "Report not found",
			Success: false,
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, BaseResponseError)
	}

	// Berhasil mendapatkan data paket
	BaseResponseSuccess := models.SuccessResponse{
		Message: "Success get report",
		Success: true,
		Data:    report,
	}
	return c.JSON(http.StatusOK, BaseResponseSuccess)
}
