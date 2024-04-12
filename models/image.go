package models

type Image struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Src      string `json:"src"`
	ReportID int    `json:"report_id"`
}

type ImageResponse struct {
	ID  int    `json:"id" gorm:"primary_key"`
	Src string `json:"src"`
}
