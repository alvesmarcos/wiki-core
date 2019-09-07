package models

import "github.com/jinzhu/gorm"

// Object - or Files
type Object struct {
	// extends
	Model
	// fields
	Name           string     `gorm:"not_null" json:"name"`
	Description    string     `json:"description"`
	Path           string     `json:"path"`
	NegativeReview uint       `gorm:"default:0" json:"negative_review"`
	PositiveReview uint       `gorm:"default:0" json:"positive_review"`
	ObjectType     ObjectType `gorm:"foreignKey:ObjectTypeID;not_null" json:"type"`
	Context        Context    `gorm:"foreignKey:ContextID;not_null" json:"context"`
	// relationships
	ObjectTypeID uint `json:"-"`
	ContextID    uint `json:"-"`
}

// NewObject -
func NewObject(
	name string,
	description string,
	path string,
	negativeReview uint,
	positiveReview uint,
	objectType ObjectType,
	context Context,
) *Object {
	return &Object{
		Name:           name,
		Description:    description,
		Path:           path,
		NegativeReview: negativeReview,
		PositiveReview: positiveReview,
		ObjectType:     objectType,
		Context:        context,
	}
}

// AddObjectConstraints -
func AddObjectConstraints(db *gorm.DB) {
	db.Model(&Object{}).AddForeignKey(
		"object_type_id", "object_types(id)", "SET NULL", "CASCADE",
	)
	db.Model(&Object{}).AddForeignKey(
		"context_id", "contexts(id)", "SET NULL", "CASCADE",
	)
}
