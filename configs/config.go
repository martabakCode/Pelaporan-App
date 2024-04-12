package configs

import (
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDB() {
	dsn := "root:@tcp(localhost:3306)/getasan?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	AutoMigrate()
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Report{}, &models.Image{}, &models.Status{}, &models.User{}, &models.Camat{})
}
