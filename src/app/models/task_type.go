package models

// TaskType - model that represents types of a task
type TaskType struct {
	Model
	Name        string `gorm:"not_null" json:"name"`
	Description string `json:"description"`
	Slug        string `gorm:"unique;not_null" json:"slug"`
	// relationships
	WorkflowID uint
}

// NewTaskType - to create a new task type
func NewTaskType(name, description, slug string) *TaskType {
	return &TaskType{Name: name, Description: description, Slug: slug}
}
