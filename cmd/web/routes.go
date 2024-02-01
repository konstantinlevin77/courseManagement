package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/konstantinlevin77/courseManagement/internal/handlers"
	"net/http"
)

func NewRouter() http.Handler {

	router := chi.NewRouter()
	router.Get("/students/getById/{id}", handlers.GetStudentByIdHandler)
	router.Get("/courses/getById/{id}", handlers.GetCourseByIdHandler)
	router.Get("/courses/getByStudentId/{id}", handlers.GetCoursesByStudentIdHandler)

	return router
}
