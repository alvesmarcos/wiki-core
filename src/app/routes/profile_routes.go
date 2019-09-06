package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// ProfileRoutes - Handler of Profiles to Sequests
type ProfileRoutes struct {
	routes []Route
}

// NewProfileRoutes -
func NewProfileRoutes(db *gorm.DB) *ProfileRoutes {
	profile := controllers.NewProfileController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/profiles",
			Handler: middlewares.Authentication(profile.IndexProfiles),
		},
	}
	return &ProfileRoutes{routes: r}
}

// GetRoutes - All routes presents in profile
func (s *ProfileRoutes) GetRoutes() []Route {
	return s.routes
}
