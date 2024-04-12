package controllers

import (
	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/helpers"
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"github.com/OctavianoRyan25/Pelaporan-App/validation"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	user := models.User{}
	user.Email = c.FormValue("email")
	user.Name = c.FormValue("name")
	user.Password = c.FormValue("password")

	err := validation.ValidateRegister(user)
	if err != nil {
		BaseResponseError := models.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Code:    400,
		}
		return c.JSON(400, BaseResponseError)
	}

	user.Password = helpers.HashPass(user.Password)

	err = configs.DB.Create(&user).Error

	if err != nil {
		BaseResponseError := models.ErrorResponse{
			Message: "Failed to register",
			Success: false,
			Code:    400,
		}
		return c.JSON(400, BaseResponseError)
	}

	BaseResponseSuccess := models.SuccessResponse{
		Message: "Register successfully",
		Success: true,
		Data:    user,
	}

	return c.JSON(200, BaseResponseSuccess)
}

func RegisterCamat(c echo.Context) error {
	camat := models.Camat{}
	camat.Email = c.FormValue("email")
	camat.Name = c.FormValue("name")
	camat.Password = c.FormValue("password")

	err := validation.ValidateRegister(models.User(camat))
	if err != nil {
		BaseResponseError := models.ErrorResponse{
			Message: err.Error(),
			Success: false,
			Code:    400,
		}
		return c.JSON(400, BaseResponseError)
	}

	camat.Password = helpers.HashPass(camat.Password)

	err = configs.DB.Create(&camat).Error

	if err != nil {
		BaseResponseError := models.ErrorResponse{
			Message: "Failed to register",
			Success: false,
			Code:    400,
		}
		return c.JSON(400, BaseResponseError)
	}

	BaseResponseSuccess := models.SuccessResponse{
		Message: "Register successfully",
		Success: true,
		Data:    camat,
	}

	return c.JSON(200, BaseResponseSuccess)
}
