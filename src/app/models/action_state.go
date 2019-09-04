package models

import (
	"github.com/jinzhu/gorm"
)

// ActionState - relation many2many between Action and State
type ActionState struct {
	// extends
	Model
	// fields
	Action Action `gorm:"foreignKey:ActionID"`
	State  State  `gorm:"foreignKey:StateID"`
	// relationships
	ActionID uint
	StateID  uint
}

// AddActionAssignmentConstraints -
func AddActionAssignmentConstraints(db *gorm.DB) {
	db.Model(&ActionState{}).AddForeignKey(
		"action_id", "actions(id)", "CASCADE", "CASCADE",
	)
	db.Model(&ActionState{}).AddForeignKey(
		"state_id", "states(id)", "CASCADE", "CASCADE",
	)
}
