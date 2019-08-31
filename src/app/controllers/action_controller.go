package controllers

import (
	"net/http"

	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// ActionController - Handle of action table
type ActionController struct {
	db *gorm.DB
}

// NewActionController - Create a ActionController
func NewActionController(db *gorm.DB) *ActionController {
	return &ActionController{db: db}
}

// IndexActions - Get all actions presents in action
func (s *ActionController) IndexActions(w http.ResponseWriter, r *http.Request) {
	var actions []models.Action

	if err := s.db.Find(&actions).Error; err != nil {
		utils.SendJSON(
			w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
		)
		return
	}
	utils.SendJSON(w, &actions, http.StatusOK)
}
