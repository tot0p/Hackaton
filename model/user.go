package model

import (
	"hackaton/utils/hash"
)

// Role constants
const (
	RoleAdmin = iota
	RoleElected
	RoleUser
)

// User struct
type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
}

// ComparePassword compares the password with the hashed password
func (u *User) ComparePassword(password string) bool {
	return hash.CheckPasswordHash(password, u.Password)
}

// RemovePassword removes the password from the user struct (for security reasons)
func (u *User) RemovePassword() {
	u.Password = ""
}
