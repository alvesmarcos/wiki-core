package models

// Orientation -
type Orientation struct {
	Model
	// fields
	Name        string `gorm:"not_null" json:"name"`
	Description string `json:"description"`
	Slug        string `gorm:"unique;not_null" json:"slug"`
}

// NewOrientation - to create a new orientation
func NewOrientation(name, description, slug string) *Orientation {
	return &Orientation{Name: name, Description: description, Slug: slug}
}
