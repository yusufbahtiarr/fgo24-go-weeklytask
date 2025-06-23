package services

type UserService interface {
	AddUser(username, password string) User
	LoginUser(username, password string) bool
}

type User struct {
	Username string
	Password string
}

type userService struct {
	users []User
}

func NewUserService() UserService {
	return &userService{
		users: []User{},
	}
}

var UserServices UserService = NewUserService()

func (u *userService) AddUser(username, password string) User {
	newUser := User{
		Username: username,
		Password: password,
	}
	u.users = append(u.users, newUser)
	return newUser
}

func (u *userService) LoginUser(username, password string) bool {
	for _, user := range u.users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
