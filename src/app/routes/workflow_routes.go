package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// WorkflowRoutes - Handler of Workflows to Sequests
type WorkflowRoutes struct {
	routes []Route
}

// NewWorkflowRoutes -
func NewWorkflowRoutes(db *gorm.DB) *WorkflowRoutes {
	workflow := controllers.NewWorkflowController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/workflows",
			Handler: middlewares.Authentication(workflow.IndexWorkflows),
		},
	}
	return &WorkflowRoutes{routes: r}
}

// GetRoutes - All routes presents in workflow
func (s *WorkflowRoutes) GetRoutes() []Route {
	return s.routes
}
