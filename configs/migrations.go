package configs

import (
	"github.com/OctavianoRyan25/Pelaporan-App/models"
	"gorm.io/gorm"
)

func AutoMigrate() {
	DB.AutoMigrate(&models.Report{}, &models.Image{}, &models.Status{}, &models.User{}, &models.Camat{})

	for _, status := range models.DefaultStatuses {
		var existingStatus models.Status
		if err := DB.Where("id = ?", status.ID).First(&existingStatus).Error; err != nil {
			if gorm.ErrRecordNotFound.Error() == err.Error() {
				// Baris dengan ID tidak ditemukan, buat baru
				DB.Create(&status)
			} else {
				panic(err)
			}
		}
	}
}
