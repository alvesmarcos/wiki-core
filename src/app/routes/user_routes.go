package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"

	"github.com/jinzhu/gorm"
)

// UserRoutes - Handler of Users to Requests
type UserRoutes struct {
	routes []Route
}

// NewUserRoutes -
func NewUserRoutes(db *gorm.DB) *UserRoutes {
	user := controllers.NewUserController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: user.IndexUsers,
		},
		Route{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: user.StoreUsers,
		},
	}
	return &UserRoutes{routes: r}
}

// GetRoutes - All routes presents in state
func (s *UserRoutes) GetRoutes() []Route {
	return s.routes
}
