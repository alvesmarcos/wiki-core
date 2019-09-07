package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ObjectStoreValidaor -
type ObjectStoreValidaor struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ContextID    uint   `json:"context_id"`
	ObjectTypeID uint   `json:"object_type_id"`
}

// Validate - ValidateContextStore
func (o ObjectStoreValidaor) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Name, validation.Required),
		validation.Field(&o.Description, validation.Required),
		validation.Field(&o.ContextID, validation.Required),
		validation.Field(&o.ObjectTypeID, validation.Required),
	)
}
