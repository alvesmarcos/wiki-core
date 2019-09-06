package models

// Profile -
type Profile struct {
	Model
	// fields
	Name        string `gorm:"not_null" json:"name"`
	Description string `json:"description"`
	Active      bool   `gorm:"default:true;not_null" json:"active"`
	Slug        string `gorm:"unique;not_null" json:"slug"`
}

// NewProfile - to create a new profile
func NewProfile(name, description, slug string) *Profile {
	return &Profile{Name: name, Description: description, Slug: slug}
}
