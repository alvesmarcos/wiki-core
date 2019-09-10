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

// ContextController - Handler of context table
type ContextController struct {
	db *gorm.DB
}

// NewContextController -
func NewContextController(db *gorm.DB) *ContextController {
	return &ContextController{db: db}
}

// IndexContexts - Get all contexts presents in context
func (c *ContextController) IndexContexts(w http.ResponseWriter, r *http.Request) {
	var contexts []models.Context

	err := c.db.Find(&contexts).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}
	utils.SendJSON(w, &contexts, http.StatusOK)
}

// StoreContexts - Create a new context
func (c *ContextController) StoreContexts(w http.ResponseWriter, r *http.Request) {
	var contextValidator validators.ContextStoreValidator

	json.NewDecoder(r.Body).Decode(&contextValidator)

	err := contextValidator.Validate()
	if exceptions.HandlerErrors(
		err, w, err, http.StatusBadRequest,
	) {
		return
	}
	var user models.User

	// make sure that your route has the middleware for auth
	claims, err := middlewares.GetUserAuth(r)

	c.db.First(&user, uint(claims["user_id"].(float64)))

	if exceptions.HandlerErrors(
		err, w, "we could not retrieve userid from token", http.StatusBadRequest,
	) {
		return
	}

	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	context := models.NewContext(
		contextValidator.Name,
		contextValidator.IsNational,
		contextValidator.Topic,
		contextValidator.Meaning,
		contextValidator.WordClass,
		false,
	)

	err = tx.Create(&context).Error
	if exceptions.HandlerErrors(
		err, w, "Duplicate key value violates unique constraint", http.StatusConflict,
	) {
		tx.Rollback()
		return
	}
	var state models.State

	c.db.Find(&models.State{Slug: "aguardando-video-referencia"}).First(&state)

	task := models.NewTask("Criar video de referencia", user, nil, *context, state, nil)

	err = tx.Create(&task).Error
	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError,
	) {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError,
	) {
		tx.Rollback()
		return
	}

	utils.SendJSON(w, context, http.StatusOK)
}
