package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// ContextRoutes - Handler of Contexts to Requests
type ContextRoutes struct {
	routes []Route
}

// NewContextRoutes -
func NewContextRoutes(db *gorm.DB) *ContextRoutes {
	context := controllers.NewContextController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/contexts",
			Handler: middlewares.Authentication(context.IndexContexts),
		},
		Route{
			Method:  http.MethodPost,
			Path:    "/contexts",
			Handler: middlewares.Authentication(context.StoreContexts),
		},
	}
	return &ContextRoutes{routes: r}
}

// GetRoutes - All routes presents in state
func (s *ContextRoutes) GetRoutes() []Route {
	return s.routes
}
