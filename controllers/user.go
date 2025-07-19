package controllers

import (
	"AuthInGo/services"
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


func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("User creation endpoint done"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")
	uc.UserService.LoginUser()
	w.Write([]byte("User login endpoint done"))
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



