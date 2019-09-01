package controllers

import (
	"net/http"

	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"

	"github.com/jinzhu/gorm"
)

// OrientationController - Handle of orientation table
type OrientationController struct {
	db *gorm.DB
}

// NewOrientationController - Create a OrientationController
func NewOrientationController(db *gorm.DB) *OrientationController {
	return &OrientationController{db: db}
}

// IndexOrientations - Get all orientations presents in orientation
func (s *OrientationController) IndexOrientations(w http.ResponseWriter, r *http.Request) {
	var orientations []models.Orientation

	err := s.db.Find(&orientations).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}
	utils.SendJSON(w, &orientations, http.StatusOK)
}
