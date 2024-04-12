package models

import "time"

type Report struct {
	ID            int        `json:"id" gorm:"primary_key"`
	Nama          string     `json:"nama"`
	NoTelepon     string     `json:"no_telepon"`
	Aduan         string     `json:"aduan"`
	Lokasi        string     `json:"lokasi"`
	CatatanLokasi string     `json:"catatan_lokasi"`
	StatusID      int        `json:"status_id"`
	Status        Status     `json:"status"`
	Images        []Image    `json:"image"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type ReportResponse struct {
	ID            int             `json:"id" gorm:"primary_key"`
	Nama          string          `json:"nama"`
	NoTelepon     string          `json:"no_telepon"`
	Aduan         string          `json:"aduan"`
	Lokasi        string          `json:"lokasi"`
	CatatanLokasi string          `json:"catatan_lokasi"`
	StatusID      int             `json:"status_id"`
	Images        []ImageResponse `json:"image" gorm:"foreignKey:ReportID"`
	CreatedAt     *time.Time      `json:"created_at"`
	UpdatedAt     *time.Time      `json:"updated_at"`
}

type UpdateReport struct {
	Status_id int `json:"status_id"`
}
