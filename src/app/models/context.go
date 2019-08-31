package models

// Context - Represents context of your applications
type Context struct {
	// extends
	Model
	// fields
	Name        string `gorm:"not_null" json:"name"`
	IsNational  bool   `gorm:"default:false" json:"is_national"`
	Topic       string `gorm:"not_null" json:"topic"`
	Meaning     string `gorm:"not_null" json:"meaning"`
	WordClass   string `gorm:"not_null" json:"word_class"`
	IsPublished bool   `gorm:"default:false" json:"is_published"`
}

// NewContext - to create a new Context
func NewContext(
	name string,
	isNational bool,
	topic string,
	meaning string,
	wordClass string,
	isPublished bool,
) *Context {
	return &Context{
		Name:        name,
		IsNational:  isNational,
		Topic:       topic,
		Meaning:     meaning,
		WordClass:   wordClass,
		IsPublished: isPublished,
	}
}
