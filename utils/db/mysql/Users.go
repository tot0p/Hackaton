package mysql

import (
	"github.com/google/uuid"
	"hackaton/model"
	"hackaton/utils/hash"
)

func CreateUser(username, password, email, phone string, role int) (model.User, error) {
	// create a UUID
	userUUID := uuid.New().String()
	// hash the password
	hashedPass, err := hash.HashPassword(password)
	if err != nil {
		return model.User{}, err
	}
	newUser := model.User{
		UUID:     userUUID,
		Username: username,
		Password: hashedPass,
		Email:    email,
		Phone:    phone,
		Role:     role,
	}
	// insert the user into the database
	_, err = DB.Conn.Exec("INSERT INTO 'users' ('uuid','username','password','email','phone','role') Values (?,?,?,?,?,?)", newUser.UUID, newUser.Username, newUser.Password, newUser.Email, newUser.Phone, newUser.Role)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}
