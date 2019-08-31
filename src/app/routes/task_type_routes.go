package routes

import (
	"net/http"
	"wikilibras-core/src/app/controllers"

	"github.com/jinzhu/gorm"
)

// TaskTypeRoutes - Handler of TaskTypes to Sequests
type TaskTypeRoutes struct {
	routes []Route
}

// NewTaskTypeRoutes -
func NewTaskTypeRoutes(db *gorm.DB) *TaskTypeRoutes {
	taskType := controllers.NewTaskTypeController(db)

	r := []Route{
		Route{
			Method:  http.MethodGet,
			Path:    "/task_types",
			Handler: taskType.IndexTaskTypes,
		},
	}
	return &TaskTypeRoutes{routes: r}
}

// GetRoutes - All routes presents in taskType
func (s *TaskTypeRoutes) GetRoutes() []Route {
	return s.routes
}
