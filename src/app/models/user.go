package models

// User -
type User struct {
	// extends
	Model
	// fields
	Name     string `gorm:"not_null" json:"name"`
	CPF      string `gorm:"unique;not_null" json:"cpf"`
	Email    string `json:"email"`
	Password string `gorm:"not_null" json:"-"`
}

// NewUser -
func NewUser(name, cpf, email, password string) *User {
	return &User{Name: name, CPF: cpf, Email: email, Password: password}
}

// UserHashPassword - encryp password
func (u *User) UserHashPassword() {
	// TODO
}
