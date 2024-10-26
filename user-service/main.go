package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jocarise/user-service/src/handlers"
	"github.com/jocarise/user-service/src/repositories"
	"github.com/jocarise/user-service/src/services"
)

const API_VERSION = "/v1"

func main() {
	userRepo := repositories.NewInMemoryUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post(API_VERSION+"/users/register", userHandler.RegisterUser)

	http.ListenAndServe(":4000", r)
}
