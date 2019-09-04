package controllers

import (
	"net/http"

	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// TaskController - Handle of task table
type TaskController struct {
	db *gorm.DB
}

// NewTaskController - Create a TaskController
func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{db: db}
}

// IndexTasks - Get all tasks presents in task
func (t *TaskController) IndexTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	err := t.db.Find(&tasks).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}

	var currentUser *models.User
	var createdByUser models.User
	var state models.State
	var context models.Context

	// index fields related
	for index, element := range tasks {
		t.db.First(&currentUser, element.CurrentUserID)
		t.db.First(&createdByUser, element.CreatedBy)
		t.db.First(&state, element.StateID)
		t.db.First(&context, element.ContextID)

		tasks[index].CurrentUser = currentUser
		tasks[index].CreatedBy = createdByUser
		tasks[index].State = state
		tasks[index].Context = context
	}

	utils.SendJSON(w, &tasks, http.StatusOK)
}
