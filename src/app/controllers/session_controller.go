package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"wikilibras-core/src/app/models"

	"wikilibras-core/src/app/exceptions"
	"wikilibras-core/src/app/utils"
	"wikilibras-core/src/app/validators"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// SessionController - Handle of session table
type SessionController struct {
	db *gorm.DB
}

// NewSessionController - Create a SessionController
func NewSessionController(db *gorm.DB) *SessionController {
	return &SessionController{db: db}
}

// StoreSession -
func (s *SessionController) StoreSession(w http.ResponseWriter, r *http.Request) {
	var sessionValidator validators.SessionStoreValidator

	json.NewDecoder(r.Body).Decode(&sessionValidator)

	err := sessionValidator.Validate()
	if exceptions.HandlerErrors(
		err, w, err, http.StatusBadRequest,
	) {
		return
	}

	var user models.User
	err = s.db.Where(&models.User{CPF: sessionValidator.CPF}).First(&user).Error
	if exceptions.HandlerErrors(
		err, w, "User not found", http.StatusNotFound,
	) {
		return
	}

	err = user.CheckPassword(sessionValidator.Password)
	if exceptions.HandlerErrors(
		err, w, "Password does not match", http.StatusNotFound,
	) {
		return
	}

	tokenStr := createToken(user.ID, w)
	token := models.NewToken(tokenStr, "bearer", false, user)

	err = s.db.Create(&token).Error
	if exceptions.HandlerErrors(
		err, w, "Duplicate key value violates unique constraint", http.StatusConflict,
	) {
		return
	}
	utils.SendJSON(w, token, http.StatusOK)
}

func createToken(userID uint, w http.ResponseWriter) string {
	apiSecret := os.Getenv("API_SECRET")
	expirationTime := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(apiSecret))

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return tokenString
}
