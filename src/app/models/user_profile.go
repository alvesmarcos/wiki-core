package models

import (
	"github.com/jinzhu/gorm"
)

// UserProfile - relation many2many between User and Profile
type UserProfile struct {
	// extends
	Model
	// fields
	User    User    `gorm:"foreignKey:UserID"`
	Profile Profile `gorm:"foreignKey:ProfileID"`
	// relationships
	UserID    uint
	ProfileID uint
}

// AddUserProfileConstraints -
func AddUserProfileConstraints(db *gorm.DB) {
	db.Model(&UserProfile{}).AddForeignKey(
		"user_id", "users(id)", "CASCADE", "CASCADE",
	)
	db.Model(&UserProfile{}).AddForeignKey(
		"profile_id", "profiles(id)", "CASCADE", "CASCADE",
	)
}
