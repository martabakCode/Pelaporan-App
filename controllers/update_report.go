package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
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
	var report models.Report
	errGet := configs.DB.Model(&models.Report{}).Where("id = ?", conv).First(&report).Error
	if errGet != nil {
		// Handle error
	}
	// URL to which the HTTP POST request will be sent
	url := "https://7103.api.greenapi.com/waInstance7103927378/sendMessage/e62cdabcaa7d48a589921524c5da0671e0cca61a16704b6890"

	// Data to be sent in the request body as JSON
	data := map[string]interface{}{
		"chatId":  report.NoTelepon + "@c.us",                                 // WhatsApp chat ID
		"message": "Laporan telah selesai. Terima kasih atas kesabaran Anda.", // Message content
	}

	// Encode data as JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.JSON(500, err)
	}

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return c.JSON(500, err)
	}

	// Set Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(500, err)
	}
	defer resp.Body.Close()

	// Read the response body
	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return c.JSON(500, err)
	}

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
