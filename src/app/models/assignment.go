package models

import "github.com/jinzhu/gorm"

// Assignment -
type Assignment struct {
	// extends
	Model
	// fields
	Details     string      `json:"details"`
	Success     bool        `gorm:"not_null;default:false" json:"success"`
	Task        Task        `gorm:"foreignKey:TaskID;not_null" json:"task"`
	User        User        `gorm:"foreignKey:UserID;not_null" json:"user"`
	Orientation Orientation `gorm:"foreignKey:OrientationID;not_null" json:"orientation"`
	Actions     []Action    `json:"actions"`
	// relationships
	TaskID        uint `json:"-"`
	UserID        uint `json:"-"`
	OrientationID uint `json:"-"`
}

// NewAssignment -
func NewAssignment(
	details string,
	success bool,
	actions []Action,
	task Task,
	user User,
	orientation Orientation,
) *Assignment {
	return &Assignment{
		Details:     details,
		Success:     success,
		Actions:     actions,
		Task:        task,
		User:        user,
		Orientation: orientation,
	}
}

// AddAssignmentConstraints -
func AddAssignmentConstraints(db *gorm.DB) {
	db.Model(&Assignment{}).AddForeignKey(
		"user_id", "users(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Assignment{}).AddForeignKey(
		"task_id", "tasks(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Assignment{}).AddForeignKey(
		"orientation_id", "orientations(id)", "SET NULL", "CASCADE",
	)
}

// LoadRelationships -
func (a *Assignment) LoadRelationships(db *gorm.DB) {
	var user User
	var task Task
	var actionstates []ActionState
	var actions []Action

	db.First(&user, a.UserID)
	db.First(&task, a.TaskID)

	db.Where(&ActionState{StateID: task.StateID}).Find(&actionstates)

	for _, element := range actionstates {
		element.LoadRelationships(db)

		actions = append(actions, element.Action)
	}

	a.User = user
	a.Task = task
	a.Actions = actions
}
