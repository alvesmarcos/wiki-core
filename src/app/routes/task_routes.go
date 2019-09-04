package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"
	"wikilibras-core/src/app/middlewares"

	"github.com/jinzhu/gorm"
)

// TaskRoutes - Handler of Tasks to Sequests
type TaskRoutes struct {
	routes []Route
}

// NewTaskRoutes -
func NewTaskRoutes(db *gorm.DB) *TaskRoutes {
	task := controllers.NewTaskController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: middlewares.Authentication(task.IndexTasks),
		},
	}
	return &TaskRoutes{routes: r}
}

// GetRoutes - All routes presents in task
func (s *TaskRoutes) GetRoutes() []Route {
	return s.routes
}
