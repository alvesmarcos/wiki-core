package models

// Action - represents action that one task can do
type Action struct {
	Model
	Description  string `json:"description"`
	Caption      string `json:"caption"`
	Help         string `json:"help"`
	Interactive  bool   `json:"interactive"`
	Confirmation bool   `json:"confirmation"`
	// relationships
	WorkflowID uint
}

// NewAction - to create a Action
func NewAction(
	description string,
	caption string,
	help string,
	interactive bool,
	confirmation bool) *Action {

	return &Action{
		Description:  description,
		Caption:      caption,
		Help:         help,
		Interactive:  interactive,
		Confirmation: confirmation,
	}
}
