package routes

import (
	"net/http"
)

// Route -
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Handler -
type Handler interface {
	GetRoutes() []Route
}
