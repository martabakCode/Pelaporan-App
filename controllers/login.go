package controllers

import (
	"net/http"

	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/helpers"
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"github.com/OctavianoRyan25/Pelaporan-App/validation"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	user := models.User{}
	password := ""
	user.Email = c.FormValue("email")
	user.Password = c.FormValue("password")

	err := validation.ValidateLogin(user)
	if err != nil {
		baseResponseError := models.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, baseResponseError)
	}

	password = user.Password

	err = configs.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "Login successfully",
		"id":     user.ID,
		"email":  user.Email,
		"name":   user.Name,
		"token":  token,
	})
}

func LoginCamat(c echo.Context) error {
	camat := models.Camat{}
	password := ""
	camat.Email = c.FormValue("email")
	camat.Password = c.FormValue("password")

	err := validation.ValidateLogin(models.User(camat))
	if err != nil {
		baseResponseError := models.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, baseResponseError)
	}

	password = camat.Password

	err = configs.DB.Where("email = ?", camat.Email).First(&camat).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	comparePass := helpers.ComparePass([]byte(camat.Password), []byte(password))
	if !comparePass {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
	}

	token, err := helpers.GenerateToken(camat.ID, camat.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "Login successfully",
		"id":     camat.ID,
		"email":  camat.Email,
		"name":   camat.Name,
		"token":  token,
	})
}
