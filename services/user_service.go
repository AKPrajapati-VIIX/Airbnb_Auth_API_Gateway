package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	GetUserById() error

	//==========
	// CreateUser(user *models.User) error
	// GenerateJWT(userID int, email string) (string, error)
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


