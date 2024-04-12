package controllers

import (
	"strconv"

	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"github.com/labstack/echo/v4"
)

func UpdateReportStatus(c echo.Context) error {
	id := c.Param("id")
	conv, err := strconv.Atoi(id)
	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: "Error converting id",
			Success: false,
			Code:    500,
		}
		return c.JSON(500, errorResponse)
	}
	status_id := c.FormValue("status_id")

	err = configs.DB.Model(&models.Report{}).Where("id = ?", conv).Update("status_id", status_id).Error
	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: "Failed to update report",
			Success: false,
			Code:    500,
		}
		return c.JSON(500, errorResponse)
	}

	successResponse := models.SuccessResponse{
		Message: "Success update report",
		Success: true,
	}

	return c.JSON(200, successResponse)
}
