package controllers

import (
	"encoding/json"
	"net/http"
	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/middlewares"
	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"
	"wikilibras-core/src/app/validators"

	"github.com/jinzhu/gorm"
)

// AssignmentController -
type AssignmentController struct {
	db *gorm.DB
}

// NewAssignmentController -
func NewAssignmentController(db *gorm.DB) *AssignmentController {
	return &AssignmentController{db: db}
}

// IndexAssignments - Get all assignments by user
func (a *AssignmentController) IndexAssignments(w http.ResponseWriter, r *http.Request) {
	var assignments []models.Assignment
	var user models.User

	// make sure that your route has the middleware for auth
	claims, err := middlewares.GetUserAuth(r)

	a.db.First(&user, uint(claims["user_id"].(float64)))

	if exceptions.HandlerErrors(
		err, w, "we could not retrieve userid from token", http.StatusBadRequest,
	) {
		return
	}

	err = a.db.Where(&models.Assignment{UserID: user.ID}).Find(&assignments).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}

	// index fields related
	for index := range assignments {
		assignments[index].LoadRelationships(a.db)
	}

	utils.SendJSON(w, &assignments, http.StatusOK)
}

// StoreAssignments -
func (a *AssignmentController) StoreAssignments(w http.ResponseWriter, r *http.Request) {
	var assignmentValidator validators.AssignmentStoreValidator

	json.NewDecoder(r.Body).Decode(&assignmentValidator)

	err := assignmentValidator.Validate()
	if exceptions.HandlerErrors(
		err, w, err, http.StatusBadRequest,
	) {
		return
	}

	var user models.User

	// make sure that your route has the middleware for auth
	claims, err := middlewares.GetUserAuth(r)

	a.db.First(&user, uint(claims["user_id"].(float64)))

	if exceptions.HandlerErrors(
		err, w, "we could not retrieve userid from token", http.StatusBadRequest,
	) {
		return
	}

	var orientation models.Orientation
	var task models.Task

	err = a.db.First(&orientation, assignmentValidator.OrientationID).Error

	if exceptions.HandlerErrors(
		err, w, "orientation_id does not exist", http.StatusNotFound,
	) {
		return
	}

	err = a.db.First(&task, assignmentValidator.TaskID).Error

	if exceptions.HandlerErrors(
		err, w, "task_id does not exist", http.StatusNotFound,
	) {
		return
	}

	task.CurrentUserID = &user.ID

	a.db.Save(&task)

	assignments := models.NewAssignment(
		assignmentValidator.Details, false, []models.Action{}, task, user, orientation,
	)

	err = a.db.Create(&assignments).Error
	if exceptions.HandlerErrors(
		err, w, "Duplicate key value violates unique constraint", http.StatusConflict,
	) {
		return
	}

	assignments.LoadRelationships(a.db)

	utils.SendJSON(w, assignments, http.StatusOK)
}
