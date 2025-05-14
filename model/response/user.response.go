package response

import (
	"time"

	"gorm.io/gorm"
)

type LoginResponse struct {
	Name  string
	Token string
}

type RegisterResponse struct {
	Name     string
	Email    string
	Address  string
	Password string
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey" `
	Name      string         `json:"name" validate:"required"`
	Password  string         `json:"-" validate:"required"`
	Address   string         `json:"Address" gorm:"column:password"`
	Phone     string         `json:"Phone"`
	Email     string         `json:"Email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}