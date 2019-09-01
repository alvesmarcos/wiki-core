package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"

	"github.com/jinzhu/gorm"
)

// SessionRoutes - Handler of Actions to Sequests
type SessionRoutes struct {
	routes []Route
}

// NewSessionRoutes -
func NewSessionRoutes(db *gorm.DB) *SessionRoutes {
	session := controllers.NewSessionController(db)

	r := []Route{
		Route{
			Method:  http.MethodPost,
			Path:    "/sessions",
			Handler: session.StoreSession,
		},
	}
	return &SessionRoutes{routes: r}
}

// GetRoutes - All routes presents in action
func (s *SessionRoutes) GetRoutes() []Route {
	return s.routes
}
