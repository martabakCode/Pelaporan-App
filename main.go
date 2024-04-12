package main

import (
	"github.com/OctavianoRyan25/Pelaporan-App/configs"
	"github.com/OctavianoRyan25/Pelaporan-App/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.CreateDB()
	e := echo.New()
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
