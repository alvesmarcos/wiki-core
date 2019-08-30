package models

import "time"

// Model - fields commonin all models
type Model struct {
	ID        uint      `gorm:"AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
