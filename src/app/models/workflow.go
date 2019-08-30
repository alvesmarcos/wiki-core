package models

// Workflow -
type Workflow struct {
	Model
	TaskType     TaskType `json:"task_type"`
	TaskTypePrev TaskType `json:"task_type_prev"`
	TaskTypeNext TaskType `json:"task_type_next"`
	Action       Action   `json:"action"`
}
