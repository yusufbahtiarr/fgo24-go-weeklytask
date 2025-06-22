package services

import (
	"go-booking-menu/menus"
)

func AddUser(username string, password string) menus.User {
	dataUser := menus.User{
		Username: username,
		Password: password,
	}
	menus.Users = append(menus.Users, dataUser)
	return dataUser
}

func LoginUser(username string, password string) bool {
	for _, user := range menus.Users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
