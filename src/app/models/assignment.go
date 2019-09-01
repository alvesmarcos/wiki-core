package models

import "github.com/jinzhu/gorm"

// Assignment -
type Assignment struct {
	// extends
	Model
	// fields
	Details     string      `json:"details"`
	Success     bool        `gorm:"not_null;default:false" json:"success"`
	TaskType    TaskType    `gorm:"foreignKey:TaskType;not_null" json:"task"`
	User        User        `gorm:"foreignKey:UserID;not_null" json:"user"`
	Orientation Orientation `gorm:"foreignKey:OrientationID;not_null" json:"orientation"`
	Actions     []Action    `json:"actions"`
	// relationships
	TaskTypeID    uint `json:"-"`
	UserID        uint `json:"-"`
	OrientationID uint `json:"-"`
}

// NewAssignment -
func NewAssignment(
	details string,
	success bool,
	actions []Action,
	taskType TaskType,
	user User,
	orientation Orientation,
) *Assignment {
	return &Assignment{
		Details:     details,
		Success:     success,
		Actions:     actions,
		TaskType:    taskType,
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
		"task_type_id", "task_types(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Assignment{}).AddForeignKey(
		"orientation_id", "orientations(id)", "SET NULL", "CASCADE",
	)
}
