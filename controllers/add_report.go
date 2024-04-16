package controllers

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"bytes"
	"encoding/json"
	"net/http"

	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"github.com/OctavianoRyan25/Pelaporan-App/validation"
	"github.com/labstack/echo/v4"
)

// func AddReport(c echo.Context) error {
// 	report := models.Report{}

// 	// Binding data report
// 	report.StatusID = 1
// 	report.Nama = c.FormValue("nama")
// 	report.NoTelepon = c.FormValue("no_telepon")
// 	report.Aduan = c.FormValue("aduan")
// 	report.Lokasi = c.FormValue("lokasi")
// 	report.CatatanLokasi = c.FormValue("catatan_lokasi")

// 	err := validation.ValidateReport(report)
// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: err.Error(),
// 			Success: false,
// 			Code:    http.StatusBadRequest,
// 		}
// 		return c.JSON(http.StatusBadRequest, errorResponse)
// 	}

// 	// Binding data report
// 	// err = c.Bind(&report)
// 	// if err != nil {
// 	// 	errorResponse := models.ErrorResponse{
// 	// 		Message: "Failed to bind report",
// 	// 		Success: false,
// 	// 		Code:    500,
// 	// 	}
// 	// 	return c.JSON(500, errorResponse)
// 	// }

// 	// Mengambil file image
// 	file, err := c.FormFile("images")
// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: "images is required",
// 			Success: false,
// 			Code:    http.StatusInternalServerError,
// 		}
// 		return c.JSON(http.StatusInternalServerError, errorResponse)
// 	}

// 	src, err := file.Open()

// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: "failed to read image",
// 			Success: false,
// 			Code:    http.StatusInternalServerError,
// 		}
// 		return c.JSON(http.StatusInternalServerError, errorResponse)

// 	}
// 	defer src.Close()

// 	// Generate gambar yang bersifa unik
// 	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
// 	ext := filepath.Ext(file.Filename)
// 	err = validation.ValidateExtensionImage(ext)
// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: err.Error(),
// 			Success: false,
// 			Code:    http.StatusBadRequest,
// 		}
// 		return c.JSON(http.StatusBadRequest, errorResponse)
// 	}
// 	err = validation.ValidateSizeImage(file.Size)
// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: err.Error(),
// 			Success: false,
// 			Code:    http.StatusBadRequest,
// 		}
// 		return c.JSON(http.StatusBadRequest, errorResponse)
// 	}
// 	newFilename := timestamp + ext

// 	dst, err := os.Create(filepath.Join("public", newFilename))
// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: "failed to save image",
// 			Success: false,
// 			Code:    http.StatusInternalServerError,
// 		}
// 		return c.JSON(http.StatusInternalServerError, errorResponse)
// 	}
// 	defer dst.Close()

// 	if _, err = io.Copy(dst, src); err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: "failed to copy image",
// 			Success: false,
// 			Code:    http.StatusInternalServerError,
// 		}
// 		return c.JSON(http.StatusInternalServerError, errorResponse)
// 	}

// 	report.Images.Src = newFilename

// 	// Save data report
// 	err = configs.DB.Create(&report).Error
// 	if err != nil {
// 		errorResponse := models.ErrorResponse{
// 			Message: err.Error(),
// 			Success: false,
// 			Code:    500,
// 		}
// 		return c.JSON(500, errorResponse)
// 	}

// 	//Mapping data response
// 	dataResponse := models.ReportResponse{
// 		ID:            report.ID,
// 		Nama:          report.Nama,
// 		NoTelepon:     report.NoTelepon,
// 		Aduan:         report.Aduan,
// 		Lokasi:        report.Lokasi,
// 		CatatanLokasi: report.CatatanLokasi,
// 		StatusID:      report.StatusID,
// 		Images: models.ImageResponse{
// 			ID:  report.Images.ID,
// 			Src: report.Images.Src,
// 		},
// 		CreatedAt: report.CreatedAt,
// 		UpdatedAt: report.UpdatedAt,
// 	}

// 	successResponse := models.SuccessResponse{
// 		Message: "Report saved",
// 		Success: true,
// 		Data:    dataResponse,
// 	}
// 	return c.JSON(200, successResponse)
// }

func AddReport(c echo.Context) error {
	report := models.Report{}

	// Binding data report
	report.StatusID = 1
	report.Nama = c.FormValue("nama")
	report.NoTelepon = c.FormValue("no_telepon")
	report.Aduan = c.FormValue("aduan")
	report.Lokasi = c.FormValue("lokasi")
	report.CatatanLokasi = c.FormValue("catatan_lokasi")

	err := validation.ValidateReport(report)
	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	// Mengambil file image
	form, err := c.MultipartForm()
	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: "failed to parse multipart form",
			Success: false,
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	files := form.File["images"]

	for _, file := range files {

		err = validation.ValidateSizeImage(file.Size)
		if err != nil {
			errorResponse := models.ErrorResponse{
				Message: err.Error(),
				Success: false,
				Code:    http.StatusBadRequest,
			}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		src, err := file.Open()
		if err != nil {
			errorResponse := models.ErrorResponse{
				Message: "failed to read image",
				Success: false,
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, errorResponse)
		}
		defer src.Close()

		// Generate gambar unik
		timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
		ext := filepath.Ext(file.Filename)
		err = validation.ValidateExtensionImage(ext)
		if err != nil {
			errorResponse := models.ErrorResponse{
				Message: err.Error(),
				Success: false,
				Code:    http.StatusBadRequest,
			}
			return c.JSON(http.StatusBadRequest, errorResponse)
		}

		newFilename := timestamp + ext

		dst, err := os.Create(filepath.Join("public", newFilename))
		if err != nil {
			errorResponse := models.ErrorResponse{
				Message: "failed to save image",
				Success: false,
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, errorResponse)
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			errorResponse := models.ErrorResponse{
				Message: "failed to copy image",
				Success: false,
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, errorResponse)
		}

		report.Images = append(report.Images, models.Image{Src: newFilename})
	}

	// Save data report
	err = configs.DB.Create(&report).Error
	if err != nil {
		errorResponse := models.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Code:    500,
		}
		return c.JSON(500, errorResponse)
	}
	// URL to which the HTTP POST request will be sent
	url := "https://7103.api.greenapi.com/waInstance7103927378/sendMessage/e62cdabcaa7d48a589921524c5da0671e0cca61a16704b6890"

	// Data to be sent in the request body as JSON
	data := map[string]interface{}{
		"chatId":  report.NoTelepon + "@c.us",                                               // WhatsApp chat ID
		"message": report.Nama + " telah berhasil dibuat dan sedang diproses. Terima kasih", // Message content
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
	// Mapping data response
	dataResponse := models.ReportResponse{
		ID:            report.ID,
		Nama:          report.Nama,
		NoTelepon:     report.NoTelepon,
		Aduan:         report.Aduan,
		Lokasi:        report.Lokasi,
		CatatanLokasi: report.CatatanLokasi,
		StatusID:      report.StatusID,
		CreatedAt:     report.CreatedAt,
		UpdatedAt:     report.UpdatedAt,
	}

	var imagesResponse []models.ImageResponse
	for _, image := range report.Images {
		imagesResponse = append(imagesResponse, models.ImageResponse{ID: image.ID, Src: image.Src})
	}
	dataResponse.Images = imagesResponse

	successResponse := models.SuccessResponse{
		Message: "Report saved",
		Success: true,
		Data:    dataResponse,
	}
	return c.JSON(http.StatusOK, successResponse)
}
