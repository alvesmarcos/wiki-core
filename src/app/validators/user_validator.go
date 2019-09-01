package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// UserStoreValidator -
type UserStoreValidator struct {
	Name                 string `json:"name"`
	CPF                  string `json:"cpf"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Email                string `json:"email"`
}

// Validate - ValidateUserStore
func (u UserStoreValidator) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.CPF, validation.Required),
		validation.Field(
			&u.PasswordConfirmation,
			validation.Required,
			validation.By(stringEquals(
				u.Password, "does not match with password provided",
			)),
		),
		validation.Field(&u.Email, is.Email),
	)
}
