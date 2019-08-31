package controllers

import (
	"net/http"

	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// WorkflowController - Handle of workflow table
type WorkflowController struct {
	db *gorm.DB
}

// NewWorkflowController - Create a WorkflowController
func NewWorkflowController(db *gorm.DB) *WorkflowController {
	return &WorkflowController{db: db}
}

// IndexWorkflows - Get all workflows presents in workflow
func (s *WorkflowController) IndexWorkflows(w http.ResponseWriter, r *http.Request) {
	var workflows []models.Workflow

	if err := s.db.Find(&workflows).Error; err != nil {
		utils.SendJSON(
			w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
		)
		return
	}

	var taskType models.TaskType
	var statePrev models.State
	var stateNext models.State
	var action models.Action

	// index fields related
	for index, element := range workflows {
		s.db.First(&taskType, element.TaskTypeID)
		s.db.First(&statePrev, element.StatePrevID)
		s.db.First(&stateNext, element.StateNextID)
		s.db.First(&action, element.ActionID)

		workflows[index].TaskType = taskType
		workflows[index].StateNext = stateNext
		workflows[index].StatePrev = statePrev
		workflows[index].Action = action
	}

	utils.SendJSON(w, &workflows, http.StatusOK)
}
