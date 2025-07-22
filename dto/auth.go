package dto



type LoginUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequestDTO struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}


