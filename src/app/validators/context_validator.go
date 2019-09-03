package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ContextStoreValidator -
type ContextStoreValidator struct {
	Name       string `json:"name"`
	IsNational bool   `json:"is_national"`
	Topic      string `json:"topic"`
	Meaning    string `json:"meaning"`
	WordClass  string `json:"word_class"`
}

// Validate - ValidateContextStore
func (c ContextStoreValidator) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.IsNational, validation.Required),
		validation.Field(&c.Topic, validation.Required),
		validation.Field(&c.Meaning, validation.Required),
		validation.Field(&c.WordClass, validation.Required),
	)
}
