package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/konstantinlevin77/courseManagement/internal/handlers"
	"net/http"
)

func NewRouter() http.Handler {

	router := chi.NewRouter()
	router.Get("/", handlers.HelloWorldHandler)
	router.Get("/users/getById/{id}", handlers.GetUserByIdHandler)

	return router
}
