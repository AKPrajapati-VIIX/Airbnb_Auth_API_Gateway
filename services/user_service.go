package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById() error

	//==========
	// CreateUser(user *models.User) error
	// GenerateJWT(userID int, email string) (string, error)
	CreateUser() error
	LoginUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	password := "example_password"
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.userRepository.Create(
		"username_example_1",
		"user1@example.com",
		hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) LoginUser() error {
	// use bcrypt password from mysql======
	response := utils.CheckPasswordHash("example_password", "$2a$10$WFL24moialK1iFd28RgEFeK7wJ3pp2TfLO95pNtV70k4FwttqYJnO")
	fmt.Println("Login response:", response)
	return nil
}