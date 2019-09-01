package exceptions

import (
	"net/http"
	"wikilibras-core/src/app/utils"
)

// ErrorShape -
type ErrorShape struct {
	Message interface{} `json:"message"`
}

// HandlerErrors - Check if exist errors
func HandlerErrors(
	err error, w http.ResponseWriter, message interface{}, code int,
) bool {
	if err != nil {
		errors := ErrorShape{Message: message}
		utils.SendJSON(w, errors, code)
		return true
	}
	return false
}
