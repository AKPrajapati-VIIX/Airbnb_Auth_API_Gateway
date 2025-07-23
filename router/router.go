package router

import (
	"AuthInGo/controllers"
	// "AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	// Middleware
	chiRouter.Use(middleware.Logger) // Log the request
	 
	// Built-in Chi middleware for logging requests
	// chiRouter.Use(middlewares.UserCreateValidator) // Middleware for validating requests
	// chiRouter.Use(middlewares.UserLoginValidator)  // Middleware for validating login requests
	// chiRouter.Use(middleware.Recoverer) // Recover from panics
	// chiRouter.Use(middlewares.RateLimiterMiddleware) // Recover from panics

	chiRouter.Get("/ping", controllers.PingHandler)

	// Register the user router
	//==============rever proxy to fakestoreapi.in  
	chiRouter.HandleFunc("/fakestoreservice/*", utils.ProxyToService("https://fakestoreapi.in", "/fakestoreservice"))

	UserRouter.Register(chiRouter)

	return chiRouter

}