package routes

import (
	"github.com/OctavianoRyan25/Pelaporan-App/controllers"
	"github.com/OctavianoRyan25/Pelaporan-App/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.Logger())

	e.POST("/register", controllers.Register)
	e.POST("/register-camat", controllers.RegisterCamat)

	e.POST("/login", controllers.Login)
	e.POST("/login-camat", controllers.LoginCamat)

	e.POST("/add-report", controllers.AddReport)

	e.GET("/get-report", controllers.GetReport, middlewares.Authentication())
	e.PUT("/update-report/:id", controllers.UpdateReportStatus, middlewares.Authentication())
}
