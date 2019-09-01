package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/models"
	"wikilibras-core/src/app/utils"
	"wikilibras-core/src/app/validators"

	"github.com/jinzhu/gorm"
)

// UserController - Handle of state table
type UserController struct {
	db *gorm.DB
}

// NewUserController - Create a UserController
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

// IndexUsers - Get all users presents in user
func (u *UserController) IndexUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	err := u.db.Find(&users).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}
	utils.SendJSON(w, &users, http.StatusOK)
}

// StoreUsers - Post User
func (u *UserController) StoreUsers(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity,
	) {
		return
	}

	var userValidator validators.UserStoreValidator
	err = json.Unmarshal(body, &userValidator)

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity,
	) {
		return
	}
	err = userValidator.Validate()

	if exceptions.HandlerErrors(
		err, w, err, http.StatusBadRequest,
	) {
		return
	}
	user := models.NewUser(
		userValidator.Name,
		userValidator.CPF,
		userValidator.Email,
		userValidator.Password,
	)
	user.HashPassword()

	err = u.db.Create(&user).Error
	if exceptions.HandlerErrors(
		err, w, "Duplicate key value violates unique constraint", http.StatusConflict,
	) {
		return
	}

	utils.SendJSON(w, user, http.StatusOK)
}
