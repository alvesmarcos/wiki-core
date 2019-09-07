package controllers

import (
	"encoding/json"
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

	for index, user := range users {
		var profiles []models.Profile
		var userProfiles []models.UserProfile
		var profile models.Profile

		u.db.Where(&models.UserProfile{UserID: user.ID}).Find(&userProfiles)

		for _, element := range userProfiles {
			u.db.First(&profile, element.ProfileID)
			profiles = append(profiles, profile)
		}
		users[index].Profiles = profiles
	}
	utils.SendJSON(w, &users, http.StatusOK)
}

// StoreUsers - Post User
func (u *UserController) StoreUsers(w http.ResponseWriter, r *http.Request) {
	var userValidator validators.UserStoreValidator

	json.NewDecoder(r.Body).Decode(&userValidator)

	err := userValidator.Validate()
	if exceptions.HandlerErrors(
		err, w, err, http.StatusBadRequest,
	) {
		return
	}

	var profiles []models.Profile

	for _, element := range userValidator.ProfilesID {
		var profile models.Profile

		err = u.db.First(&profile, element).Error

		if exceptions.HandlerErrors(
			err, w, "profile_id does not exist", http.StatusNotFound,
		) {
			return
		}
		profiles = append(profiles, profile)
	}

	user := models.NewUser(
		userValidator.Name,
		userValidator.CPF,
		userValidator.Email,
		userValidator.Password,
		profiles,
	)
	user.HashPassword()

	err = u.db.Create(&user).Error
	if exceptions.HandlerErrors(
		err, w, "Duplicate key value violates unique constraint", http.StatusConflict,
	) {
		return
	}

	for _, element := range profiles {
		u.db.Create(&models.UserProfile{User: *user, Profile: element})
	}

	utils.SendJSON(w, user, http.StatusOK)
}
