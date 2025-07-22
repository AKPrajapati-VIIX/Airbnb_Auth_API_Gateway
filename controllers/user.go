package controllers

import (
	"AuthInGo/services"
	"AuthInGo/utils"
	"AuthInGo/dto"
	"github.com/go-chi/chi/v5"
	"fmt"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in UserController")
	uc.UserService.GetUserById()
	w.Write([]byte("User fetching endpoint done"))
}


// 
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	

	var payload dto.CreateUserRequestDTO

	// Read and decode JSON
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while creating user", jsonErr)
		return
	}

	fmt.Println("Payload received:", payload)

	// Validate the decoded struct
	if err := utils.Validator.Struct(payload); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
		return
	}

	// Create the user
	user, err := uc.UserService.CreateUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	// Send success response
	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created successfully:", user)
}



func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("LoginUser called in UserController")
	// uc.UserService.LoginUser()
	// w.Write([]byte("User login endpoint done"))

	//==============adding some validation to think more==============
	// var payload dto.LoginUserRequestDTO
	// Read and decode JSON

	fmt.Println("LoginUser called in UserController")
	// Assuming you have a LoginUserRequestDTO defined in your dto package
	var payload dto.LoginUserRequestDTO

	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Something went wrong while logging in", jsonErr)
		return
	}

	fmt.Println("Payload received:", payload)

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid input data", validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)

	
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllUsers called in UserController")
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	// Here you would typically convert users to JSON and write to response
	fmt.Fprintf(w, "Fetched %d users successfully", len(users))
}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUserById called in UserController")
	// Assuming you have a way to get the user ID from the request
	// id := r.URL.Query().Get("id")

	// For RESTful API, we can get the ID from the URL parameters
	// Example: DELETE /users/{id}
    id := chi.URLParam(r, "id") // Get {id} from the URL
	
	// Convert id to int64
	var userID int64
	var err error
	if id != "" {
		userID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	err = uc.UserService.DeleteUserById(userID)
	
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User deletion endpoint done"))
}



