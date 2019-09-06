package controllers

import (
	"net/http"

	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// ProfileController - Handle of profile table
type ProfileController struct {
	db *gorm.DB
}

// NewProfileController - Create a ProfileController
func NewProfileController(db *gorm.DB) *ProfileController {
	return &ProfileController{db: db}
}

// IndexProfiles - Get all profiles presents in profile
func (s *ProfileController) IndexProfiles(w http.ResponseWriter, r *http.Request) {
	var profiles []models.Profile

	err := s.db.Find(&profiles).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}
	utils.SendJSON(w, &profiles, http.StatusOK)
}
