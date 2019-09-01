package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// SessionStoreValidator -
type SessionStoreValidator struct {
	CPF      string `json:"cpf"`
	Password string `json:"password"`
}

// Validate - ValidateSessionStore
func (s SessionStoreValidator) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.CPF, validation.Required),
		validation.Field(&s.CPF, validation.Required),
	)
}
