package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// SendJSON - Send response in JSON format
func SendJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	bjson, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprint("Erro while encoding JSON: ", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"erro": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(bjson))
	}
}
