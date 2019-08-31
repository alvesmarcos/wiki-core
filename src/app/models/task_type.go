package models

import (
	"github.com/jinzhu/gorm"
)

// TaskType - model that represents types of a task
type TaskType struct {
	// extends
	gorm.Model
	// fields
	Name        string `gorm:"not_null" json:"name"`
	Description string `json:"description"`
	Slug        string `gorm:"unique;not_null" json:"slug"`
}

// NewTaskType - to create a new task type
func NewTaskType(name, description, slug string) *TaskType {
	return &TaskType{Name: name, Description: description, Slug: slug}
}
