package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"

	"github.com/jinzhu/gorm"
)

// StateRoutes - Handler of States to Sequests
type StateRoutes struct {
	routes []Route
}

// NewStateRoutes -
func NewStateRoutes(db *gorm.DB) *StateRoutes {
	state := controllers.NewStateController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/states",
			Handler: state.IndexStates,
		},
	}
	return &StateRoutes{routes: r}
}

// GetRoutes - All routes presents in state
func (s *StateRoutes) GetRoutes() []Route {
	return s.routes
}
