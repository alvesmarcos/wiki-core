package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// OrientationRoutes - Handler of Orientations to Sequests
type OrientationRoutes struct {
	routes []Route
}

// NewOrientationRoutes -
func NewOrientationRoutes(db *gorm.DB) *OrientationRoutes {
	orientation := controllers.NewOrientationController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/orientations",
			Handler: middlewares.Authentication(orientation.IndexOrientations),
		},
	}
	return &OrientationRoutes{routes: r}
}

// GetRoutes - All routes presents in orientation
func (s *OrientationRoutes) GetRoutes() []Route {
	return s.routes
}
