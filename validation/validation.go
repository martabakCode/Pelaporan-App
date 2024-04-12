package validation

import (
	"errors"

	"github.com/OctavianoRyan25/Pelaporan-App/models"
)

func ValidateReport(report models.Report) error {
	if report.Nama == "" {
		return errors.New("nama is required")
	}
	if report.NoTelepon == "" {
		return errors.New("no Telepon is required")
	}
	if report.Aduan == "" {
		return errors.New("aduan is required")
	}
	if report.Lokasi == "" {
		return errors.New("lokasi is required")
	}
	if report.CatatanLokasi == "" {
		return errors.New("catatan lokasi is required")
	}
	return nil
}

func ValidateExtensionImage(ext string) error {
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return errors.New("image must be jpg, jpeg, or png")
	}
	return nil
}

func ValidateSizeImage(size int64) error {
	if size > 1024*1024*2 {
		return errors.New("image size must be less than 2MB")
	}
	return nil
}

func ValidateRegister(user models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func ValidateLogin(user models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
