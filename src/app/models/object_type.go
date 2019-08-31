package models

// ObjectType - Support files
type ObjectType struct {
	// extends
	Model
	// fields
	Description string `gorm:"not_null"`
	MineType    string `gorm:"unique;not_null"`
}

// NewObjectType - to create a new ObjectType
func NewObjectType(description, mineType string) *ObjectType {
	return &ObjectType{Description: description, MineType: mineType}
}
