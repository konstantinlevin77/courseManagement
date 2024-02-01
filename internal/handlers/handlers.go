package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/konstantinlevin77/courseManagement/internal/config"
	"log"
	"net/http"
	"strconv"
)

// GetStudentByIdHandler is the responding handler function for endpoint /students/getById/{id}, it writes
// out the student whose id is passed.
func GetStudentByIdHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("An error happened while getting URLParam")
		log.Println()
		_, _ = w.Write([]byte("{}"))
		return
	}

	st, err := config.App.Repo.GetStudentById(id)
	if err != nil {
		log.Println("An error happened while getting from DB ")
		log.Println(err)
		_, _ = w.Write([]byte("{}"))
		return
	}

	// RETURN STUDENT JSON
	studentJSON, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		_, _ = w.Write([]byte("{}"))
		return
	}

	_, _ = w.Write(studentJSON)

}

// GetCourseByIdHandler is the responding handler function for endpoint /courses/getById/{id}, it writes
// out the course whose id is passed.
func GetCourseByIdHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("An error happened while getting URLParam")
		log.Println()
		_, _ = w.Write([]byte("{}"))
		return
	}

	st, err := config.App.Repo.GetCourseById(id)
	if err != nil {
		log.Println("An error happened while getting from DB ")
		log.Println(err)
		_, _ = w.Write([]byte("{}"))
		return
	}

	// RETURN STUDENT JSON
	courseJSON, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		_, _ = w.Write([]byte("{}"))
		return
	}

	_, _ = w.Write(courseJSON)

}

// GetCoursesByStudentIdHandler is the responding handler function for the endpoint /courses/getByStudentId/{id}, it writes
// out the courses which the student whose id is passed has access.
func GetCoursesByStudentIdHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("An error happened while getting URLParam")
		log.Println()
		_, _ = w.Write([]byte("{}"))
		return
	}

	courses, err := config.App.Repo.GetCoursesByStudentId(id)
	if err != nil {
		log.Println("An error happened while getting from DB ")
		log.Println(err)
		_, _ = w.Write([]byte("{}"))
		return
	}

	// RETURN STUDENT JSON
	coursesJSON, err := json.MarshalIndent(courses, "", "  ")
	if err != nil {
		_, _ = w.Write([]byte("{}"))
		return
	}

	_, _ = w.Write(coursesJSON)

}
