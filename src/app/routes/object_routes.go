package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// ObjectRoutes - Handler of Objects to Requests
type ObjectRoutes struct {
	routes []Route
}

// NewObjectRoutes -
func NewObjectRoutes(db *gorm.DB) *ObjectRoutes {
	object := controllers.NewObjectController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/objects",
			Handler: middlewares.Authentication(object.IndexObjects),
		},
		Route{
			Method:  http.MethodPost,
			Path:    "/objects",
			Handler: middlewares.Authentication(object.StoreObjects),
		},
	}
	return &ObjectRoutes{routes: r}
}

// GetRoutes - All routes presents in state
func (s *ObjectRoutes) GetRoutes() []Route {
	return s.routes
}
