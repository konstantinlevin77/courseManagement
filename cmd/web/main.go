package main

import (
	"github.com/konstantinlevin77/courseManagement/internal/config"
	"github.com/konstantinlevin77/courseManagement/internal/repository"
	"github.com/konstantinlevin77/courseManagement/pkg/security"
	"log"
	"net/http"
)

const PORT = ":8080"
const DSN = "host=localhost port=5432 dbname=coursemanagement user=mehmettekman password="

func main() {

	testPass := "mauro_icardi"
	hashedPass, _ := security.HashPassword(testPass)
	log.Println(hashedPass)
	log.Println(security.DoPasswordsMatch(hashedPass, testPass))

	postgresRepo, err := repository.NewPostgresRepo(DSN)
	if err != nil {
		log.Fatal("A problem happened while initializing the repo ", err)
	}

	config.NewApp(&postgresRepo)

	log.Println("Listening and serving on ", PORT)
	server := &http.Server{
		Addr:    PORT,
		Handler: NewRouter(),
	}

	err = server.ListenAndServe()
	log.Fatal("Server halted, related error message: ", err)

}
