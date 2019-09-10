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
	CurrentUser *User      `gorm:"foreignKey:CurrentUserID" json:"current_user"`
	Context     Context    `gorm:"foreignKey:ContextID" json:"context"`
	State       State      `form:"foreignKey:StateID" json:"state"`
	FineshedAt  *time.Time `json:"fineshed_at"`
	// relationships
	CreatedByID   uint  `json:"-"`
	CurrentUserID *uint `json:"-"`
	ContextID     uint  `json:"-"`
	StateID       uint  `json:"-"`
}

// NewTask -
func NewTask(
	details string,
	createdBy User,
	currentUser *User,
	context Context,
	state State,
	fineshedAt *time.Time,
) *Task {
	return &Task{
		Details:     details,
		CreatedBy:   createdBy,
		CurrentUser: currentUser,
		State:       state,
		Context:     context,
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
	db.Model(&Task{}).AddForeignKey(
		"context_id", "contexts(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Task{}).AddForeignKey(
		"state_id", "states(id)", "SET NULL", "CASCADE",
	)
}

// LoadRelationships -
func (t *Task) LoadRelationships(db *gorm.DB) {
	var createdBy User
	var currentUser *User
	var state State
	var context Context

	db.First(&createdBy, t.CreatedBy)
	db.First(&state, t.StateID)
	db.First(&context, t.ContextID)

	createdBy.LoadRelationships(db)

	if t.CurrentUserID != nil {
		var user User

		db.First(&user, *t.CurrentUserID)
		user.LoadRelationships(db)

		currentUser = &user
	}

	t.CreatedBy = createdBy
	t.CurrentUser = currentUser
	t.State = state
	t.Context = context
}
