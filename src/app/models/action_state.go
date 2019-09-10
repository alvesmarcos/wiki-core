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

// AddActionStateConstraints -
func AddActionStateConstraints(db *gorm.DB) {
	db.Model(&ActionState{}).AddForeignKey(
		"action_id", "actions(id)", "CASCADE", "CASCADE",
	)
	db.Model(&ActionState{}).AddForeignKey(
		"state_id", "states(id)", "CASCADE", "CASCADE",
	)
}

// LoadRelationships -
func (t *ActionState) LoadRelationships(db *gorm.DB) {
	var state State
	var action Action

	db.First(&state, t.StateID)
	db.First(&action, t.ActionID)

	t.State = state
	t.Action = action
}
