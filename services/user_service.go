package services

// File: services/user_service.go

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	// "AuthInGo/db/models"
	"AuthInGo/dto"
	

	env "AuthInGo/config/env"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error

	//==========
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)

	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	// LoginUser() (string, error)
	//==========
	GetAllUsers() ([]*models.User, error)
	DeleteUserById(id int64) error

}


type UserServiceImpl struct {
	userRepository db.UserRepository
}

type User struct {
	Id        int64
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
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

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user in UserService")

	// Hash the password from payload

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return nil, err
	}

	// Create user model
	user, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}

	// Step 3. Return the created user
	return user, nil
}


func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	// Pre-requisite: This function will be given email and password as parameter, which we can hardcode for now
	// email := "user1@example.com"
	// password := "example_password"
	email := payload.Email 
	// password := payload.Password
	password := payload.Password
	// Step 1. Make a repository call to get the user by email
	user, err := u.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "", err
	}

	// Step 2. If user exists, or not. If not exists, return error
	if user == nil {
		fmt.Println("No user found with the given email")
		return "", fmt.Errorf("no user found with email: %s", email)
	}

	// Step 3. If user exists, check the password using utils.CheckPasswordHash
	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", nil
	}

	// Step 4. If password matches, print a JWT token, else return error saying password does not match
	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil

}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	fmt.Println("Fetching all users in UserService")
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}


// func (u *UserServiceImpl) DeleteUserById(userID int64) error {
// 	fmt.Println("Deleting user in UserService")
// 	err := u.userRepository.DeleteUserById(userID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (us *UserServiceImpl) DeleteUserById(id int64) error {
	fmt.Println("Deleting user in UserService")
	return us.userRepository.DeleteUserById(id)
}