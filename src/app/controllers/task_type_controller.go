package controllers

import (
	"net/http"

	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// TaskTypeController - Handle of taskType table
type TaskTypeController struct {
	db *gorm.DB
}

// NewTaskTypeController - Create a TaskTypeController
func NewTaskTypeController(db *gorm.DB) *TaskTypeController {
	return &TaskTypeController{db: db}
}

// IndexTaskTypes - Get all taskTypes presents in taskType
func (t *TaskTypeController) IndexTaskTypes(w http.ResponseWriter, r *http.Request) {
	var taskTypes []models.TaskType

	err := t.db.Find(&taskTypes).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}
	utils.SendJSON(w, &taskTypes, http.StatusOK)
}
