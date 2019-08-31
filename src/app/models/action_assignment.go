package models

import (
	"github.com/jinzhu/gorm"
)

// ActionAssignment - relation many2many between Action and Assignment
type ActionAssignment struct {
	// extends
	Model
	// fields
	Action     Action     `gorm:"foreignKey:ActionID"`
	Assignment Assignment `gorm:"foreignKey:AssignmentID"`
	// relationships
	ActionID     uint
	AssignmentID uint
}

// AddActionAssignmentConstraints -
func AddActionAssignmentConstraints(db *gorm.DB) {
	db.Model(&ActionAssignment{}).AddForeignKey(
		"action_id", "actions(id)", "CASCADE", "CASCADE",
	)
	db.Model(&ActionAssignment{}).AddForeignKey(
		"assignment_id", "assignments(id)", "CASCADE", "CASCADE",
	)
}
