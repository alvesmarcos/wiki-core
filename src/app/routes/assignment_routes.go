package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// AssignmentRoutes -
type AssignmentRoutes struct {
	routes []Route
}

// NewAssignmentRoutes -
func NewAssignmentRoutes(db *gorm.DB) *AssignmentRoutes {
	assignment := controllers.NewAssignmentController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/assignments",
			Handler: middlewares.Authentication(assignment.IndexAssignments),
		},
		Route{
			Method:  http.MethodPost,
			Path:    "/assignments",
			Handler: middlewares.Authentication(assignment.StoreAssignments),
		},
	}
	return &AssignmentRoutes{routes: r}
}

// GetRoutes - All routes presents in assignment
func (s *AssignmentRoutes) GetRoutes() []Route {
	return s.routes
}
