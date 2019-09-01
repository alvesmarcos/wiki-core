package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// ActionRoutes - Handler of Actions to Sequests
type ActionRoutes struct {
	routes []Route
}

// NewActionRoutes -
func NewActionRoutes(db *gorm.DB) *ActionRoutes {
	action := controllers.NewActionController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/actions",
			Handler: middlewares.Authentication(action.IndexActions),
		},
	}
	return &ActionRoutes{routes: r}
}

// GetRoutes - All routes presents in action
func (s *ActionRoutes) GetRoutes() []Route {
	return s.routes
}
