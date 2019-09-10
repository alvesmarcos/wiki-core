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

// ObjectController - Handle of object table
type ObjectController struct {
	db *gorm.DB
}

// NewObjectController - Create a ObjectController
func NewObjectController(db *gorm.DB) *ObjectController {
	return &ObjectController{db: db}
}

// IndexObjects - Get all objects presents in object
func (o *ObjectController) IndexObjects(w http.ResponseWriter, r *http.Request) {
	var objects []models.Object

	err := o.db.Find(&objects).Error

	if exceptions.HandlerErrors(
		err, w, http.StatusText(http.StatusNotFound), http.StatusNotFound,
	) {
		return
	}

	// index fields related
	for index := range objects {
		objects[index].LoadRelationships(o.db)
	}

	utils.SendJSON(w, &objects, http.StatusOK)
}

// StoreObjects -
func (o *ObjectController) StoreObjects(w http.ResponseWriter, r *http.Request) {
	var objectValidator validators.ObjectStoreValidaor

	json.NewDecoder(r.Body).Decode(&objectValidator)

	err := objectValidator.Validate()
	if exceptions.HandlerErrors(
		err, w, err, http.StatusBadRequest,
	) {
		return
	}

	var objectType models.ObjectType
	var context models.Context

	err = o.db.First(&objectType, objectValidator.ObjectTypeID).Error

	if exceptions.HandlerErrors(
		err, w, "object_type_id does not exist", http.StatusNotFound,
	) {
		return
	}

	err = o.db.First(&context, objectValidator.ContextID).Error

	if exceptions.HandlerErrors(
		err, w, "context_id does not exist", http.StatusNotFound,
	) {
		return
	}

	object := models.NewObject(
		objectValidator.Name, objectValidator.Description, "", 0, 0, objectType, context,
	)

	err = o.db.Create(&object).Error
	if exceptions.HandlerErrors(
		err, w, "Duplicate key value violates unique constraint", http.StatusConflict,
	) {
		return
	}

	utils.SendJSON(w, object, http.StatusOK)
}
