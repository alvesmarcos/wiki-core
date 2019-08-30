package models

import "time"

// Base - fields commonin all models
type Base struct {
	ID        int       `gorm:"AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
