package models

import "time"

// Model -
type Model struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
