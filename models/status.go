package models

type Status struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Status string `json:"status"`
}

var DefaultStatuses = []Status{
	{ID: 1, Status: "Diterima"},
	{ID: 2, Status: "Menunggu"},
	{ID: 3, Status: "Diproses"},
	{ID: 4, Status: "Selesai"},
	{ID: 5, Status: "Ditolak"},
}
