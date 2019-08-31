package controllers

import (
	"net/http"

	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// StateController - Handle of state table
type StateController struct {
	db *gorm.DB
}

// NewStateController - Create a StateController
func NewStateController(db *gorm.DB) *StateController {
	return &StateController{db: db}
}

// IndexStates - Get all states presents in state
func (s *StateController) IndexStates(w http.ResponseWriter, r *http.Request) {
	var states []models.State

	if err := s.db.Find(&states).Error; err != nil {
		utils.SendJSON(
			w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
		)
		return
	}
	utils.SendJSON(w, &states, http.StatusOK)
}
