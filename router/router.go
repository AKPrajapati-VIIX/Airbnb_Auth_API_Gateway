package router

import (
	"AuthInGo/controllers"

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
	chiRouter.Use(middleware.Recoverer) // Recover from panics

	chiRouter.Get("/ping", controllers.PingHandler)

	UserRouter.Register(chiRouter)

	return chiRouter

}