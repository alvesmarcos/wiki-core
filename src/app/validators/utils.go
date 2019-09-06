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

func uniqueID(erroMsg string) validation.RuleFunc {
	return func(value interface{}) error {
		ids, _ := value.([]uint)

		for i := 0; i < len(ids); i++ {
			for j := i; j < len(ids); j++ {
				if ids[i] == ids[j] {
					return errors.New(erroMsg)
				}
			}
		}
		return nil
	}
}
