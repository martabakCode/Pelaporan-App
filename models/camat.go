package models

import (
	"time"

	"gorm.io/gorm"
)

type Camat struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at"`
}

type CamatResponse struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeleteAt  *time.Time `json:"delete_at"`
}

type CamatLogin struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
