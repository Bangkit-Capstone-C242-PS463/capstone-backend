package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string  `gorm:"primaryKey" json:"id"`
	Email        string  `gorm:"unique;not null" json:"email"`
	Password     *string `json:"password"`
	PasswordSalt *[]byte `json:"password_salt"`
	Name         string  `json:"name"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
