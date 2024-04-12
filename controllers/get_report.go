package controllers

import (
	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"github.com/labstack/echo/v4"
)

func GetReport(c echo.Context) error {
	reports := []models.Report{}
	result := configs.DB.Preload("Images").Preload("Status").Find(&reports)
	if result.Error != nil {
		errorResponse := models.ErrorResponse{
			Message: "Failed to get report",
			Success: false,
			Code:    500,
		}
		return c.JSON(500, errorResponse)
	}
	successResponse := models.SuccessResponse{
		Message: "Success get report",
		Success: true,
		Data:    reports,
	}
	return c.JSON(200, successResponse)
}
