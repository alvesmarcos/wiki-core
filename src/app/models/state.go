package models

// State - model that represents states of one task
type State struct {
	// extends
	Model
	// fields
	Name        string `gorm:"not_null" json:"name"`
	Description string `json:"description"`
	Slug        string `gorm:"unique;not_null" json:"slug"`
}

// NewState - to create a new state
func NewState(name, description, slug string) *State {
	return &State{Name: name, Description: description, Slug: slug}
}
