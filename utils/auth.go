package utils

import (
    "golang.org/x/crypto/bcrypt"
)

type LoginUserRequestDTO struct {
Email    string `json:"email" validate:"required,email"`
Password string `json:"password" validate:"required,min=8"`
}

// {"email", "password"}
type CreateUserRequestDTO struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// HashPassword hashes a plain password using bcrypt.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}