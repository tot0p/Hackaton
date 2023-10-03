package mysql

import (
	"errors"
	"github.com/google/uuid"
	"hackaton/model"
	"hackaton/utils/hash"
)

func CreateUser(username, password, email, phone string, role int) (model.User, error) {
	// check if the username is already taken
	u, _ := GetUserByUsername(username)
	if u.Username != "" {
		// username already taken
		return model.User{}, errors.New("username already taken")
	}
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
	_, err = DB.Conn.Exec("INSERT INTO users (uuid, username, password, email, phone, role) VALUES (?, ?, ?, ?, ?, ?)", newUser.UUID, newUser.Username, newUser.Password, newUser.Email, newUser.Phone, newUser.Role)
	if err != nil {
		return model.User{}, err
	}
	return newUser, nil
}

func GetUserByUUID(uuid string) (model.User, error) {
	user := model.User{}
	val := DB.Conn.QueryRow("SELECT * FROM users WHERE uuid = ?", uuid)
	err := val.Scan(&user.UUID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Role)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	val := DB.Conn.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err := val.Scan(&user.UUID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Role)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func UpdateUser(user model.User) error {
	_, err := DB.Conn.Exec("UPDATE users SET username = ?, password = ?, email = ?, phone = ?, role = ? WHERE uuid = ?", user.Username, user.Password, user.Email, user.Phone, user.Role, user.UUID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(uuid string) error {
	_, err := DB.Conn.Exec("DELETE FROM users WHERE uuid = ?", uuid)
	if err != nil {
		return err
	}
	return nil
}
