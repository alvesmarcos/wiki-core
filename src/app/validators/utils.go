package validators

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func stringEquals(str, erroMsg string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != str {
			return errors.New(erroMsg)
		}
		return nil
	}
}
