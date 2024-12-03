package dto

import (
	"time"
)

type GetAllHistoryResponse struct {
	History []History `json:"history"`
}

type History struct {
	ID     string `gorm:"primaryKey" json:"id"`
	UserID string `json:"user_id"`
	Result string `json:"result"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
