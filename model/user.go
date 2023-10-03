package model

const (
	RoleAdmin = iota
	RoleElected
	RoleUser
)

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`
}

func (u *User) RemovePassword() {
	u.Password = ""
}
