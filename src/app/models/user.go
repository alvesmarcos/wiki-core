package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User -
type User struct {
	// extends
	Model
	// fields
	Name     string    `gorm:"not_null" json:"name"`
	CPF      string    `gorm:"unique;not_null" json:"cpf"`
	Email    string    `json:"email"`
	Password string    `gorm:"not_null" json:"-"`
	Profiles []Profile `json:"profiles"`
}

// NewUser -
func NewUser(name string, cpf string, email string, password string, profiles []Profile) *User {
	return &User{Name: name, CPF: cpf, Email: email, Password: password, Profiles: profiles}
}

// HashPassword - encryp password
func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword - decode hash
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
