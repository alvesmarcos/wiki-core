package models

import (
	"github.com/jinzhu/gorm"
)

// Token -
type Token struct {
	// extends
	Model
	// fields
	Token     string `gorm:"unique;not_null" json:"token"`
	Type      string `gorm:"not_null" json:"type"`
	IsRevoked bool   `gorm:"default:false" json:"is_revoked"`
	User      User   `gorm:"foreignKey:UserID" json:"usuario"`
	// relationships
	UserID uint `json:"-"`
}

// NewToken -
func NewToken(token string, _type string, isRevoked bool, user User) *Token {
	return &Token{Token: token, Type: _type, IsRevoked: isRevoked, User: user}
}

// AddTokenConstraints -
func AddTokenConstraints(db *gorm.DB) {
	db.Model(&Token{}).AddForeignKey(
		"user_id", "users(id)", "CASCADE", "CASCADE",
	)
}
