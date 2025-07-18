package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Get("/profile", ur.userController.GetUserById)
	r.Post("/signup", ur.userController.CreateUser)
	r.Post("/login", ur.userController.LoginUser)
	// Register the route for getting all users
	r.Get("/users", ur.userController.GetAllUsers)
	// DELETE user 
	// r.Post("/delete", ur.userController.DeleteUserById)
	
	//===========  RESTFULL API for deleting a user by ID
	r.Get("/user/delete/{id}", ur.userController.DeleteUserById)


}
