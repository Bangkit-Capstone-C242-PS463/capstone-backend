package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string     `gorm:"primaryKey" json:"id"`
	Username     string     `gorm:"unique;not null" json:"username"`
	Password     *string    `json:"password"`
	PasswordSalt *[]byte    `json:"password_salt"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
