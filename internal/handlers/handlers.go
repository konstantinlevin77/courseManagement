package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/konstantinlevin77/courseManagement/internal/config"
	"log"
	"net/http"
	"strconv"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {

	message := "Hello, world!"

	_, _ = w.Write([]byte(message))

}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

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
