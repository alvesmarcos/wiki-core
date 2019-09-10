package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// AssignmentStoreValidator -
type AssignmentStoreValidator struct {
	Details       string `json:"details"`
	TaskID        uint   `json:"task_id"`
	OrientationID uint   `json:"orientation_id"`
}

// Validate - ValidateContextStore
func (a AssignmentStoreValidator) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Details, validation.Required),
		validation.Field(&a.TaskID, validation.Required),
		validation.Field(&a.OrientationID, validation.Required),
	)
}
