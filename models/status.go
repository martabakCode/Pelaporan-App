package models

type Status struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Status string `json:"status"`
}
