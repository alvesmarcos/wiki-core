package models

import (
	"github.com/jinzhu/gorm"
)

// Action - represents action that one task can do
type Action struct {
	// extends
	gorm.Model
	// fields
	Description  string `gorm:"not_null" json:"description"`
	Caption      string `gorm:"not_null" json:"caption"`
	Help         string `json:"help"`
	Interactive  bool   `gorm:"not_null" json:"interactive"`
	Confirmation bool   `gorm:"not_null" json:"confirmation"`
	Slug         string `gorm:"unique;not_null" json:"slug"`
}

// NewAction - to create a Action
func NewAction(
	description string,
	caption string,
	help string,
	interactive bool,
	confirmation bool,
	slug string,
) *Action {

	return &Action{
		Description:  description,
		Caption:      caption,
		Help:         help,
		Interactive:  interactive,
		Confirmation: confirmation,
		Slug:         slug,
	}
}
