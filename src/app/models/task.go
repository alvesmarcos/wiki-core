package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task -
type Task struct {
	// extends
	Model
	// fields
	Details     string     `gorm:"not_null" json:"details"`
	CreatedBy   User       `gorm:"foreignkey:CreatedByID;not_null" json:"created_by"`
	CurrentUser User       `gorm:"foreignKey:CurrentUserID" json:"current_user"`
	FineshedAt  *time.Time `json:"fineshed_at"`
	// relationships
	CreatedByID   uint `json:"-"`
	CurrentUserID uint `json:"-"`
}

// NewTask -
func NewTask(
	details string,
	createdBy User,
	currentUser User,
	fineshedAt *time.Time,
) *Task {
	return &Task{
		Details:     details,
		CreatedBy:   createdBy,
		CurrentUser: currentUser,
		FineshedAt:  fineshedAt,
	}
}

// AddTaskConstraints -
func AddTaskConstraints(db *gorm.DB) {
	db.Model(&Task{}).AddForeignKey(
		"created_by_id", "users(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Task{}).AddForeignKey(
		"current_user_id", "users(id)", "SET NULL", "CASCADE",
	)
}
