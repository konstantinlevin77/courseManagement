package main

import (
	"github.com/konstantinlevin77/courseManagement/internal/config"
	"github.com/konstantinlevin77/courseManagement/internal/repository"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {

	postgresRepo := repository.PostgresRepo{}

	config.NewApp(&postgresRepo)

	log.Println("Listening and serving on ", PORT)
	server := &http.Server{
		Addr:    PORT,
		Handler: NewRouter(),
	}

	err := server.ListenAndServe()
	log.Fatal("Server halted, related error message: ", err)

}
