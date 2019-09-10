package controllers

import (
	"net/http"

	"wikilibras-core/src/app/exceptions"
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
func (wf *WorkflowController) IndexWorkflows(w http.ResponseWriter, r *http.Request) {
	var workflows []models.Workflow

	err := wf.db.Find(&workflows).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}

	for index := range workflows {
		workflows[index].LoadRelationships(wf.db)
	}

	utils.SendJSON(w, &workflows, http.StatusOK)
}
