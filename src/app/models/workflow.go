package models

import (
	"github.com/jinzhu/gorm"
)

// Workflow - A mapper of all states that a task can be
type Workflow struct {
	// extends
	gorm.Model
	// fields
	TaskType  TaskType `gorm:"foreignkey:TaskTypeID" json:"task_type"`
	StatePrev State    `gorm:"foreignkey:StatePrevID" json:"state_prev"`
	StateNext State    `gorm:"foreignkey:StateNextID" json:"state_next"`
	Action    Action   `gorm:"foreignkey:ActionID" json:"action"`
	// relationships
	TaskTypeID  uint
	StatePrevID uint
	StateNextID uint
	ActionID    uint
}

// AddWorkflowConstraints -
func AddWorkflowConstraints(db *gorm.DB) {
	db.Model(&Workflow{}).AddForeignKey(
		"task_type_id", "task_types(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Workflow{}).AddForeignKey(
		"state_prev_id", "states(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Workflow{}).AddForeignKey(
		"state_next_id", "states(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Workflow{}).AddForeignKey(
		"action_id", "actions(id)", "SET NULL", "CASCADE",
	)
}

// NewWorkflow - to create a NewWorkflow
func NewWorkflow(
	taskType TaskType,
	statePrev State,
	stateNext State,
	action Action,
) *Workflow {
	return &Workflow{
		TaskType:  taskType,
		StatePrev: statePrev,
		StateNext: stateNext,
		Action:    action,
	}
}
