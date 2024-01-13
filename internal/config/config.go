package config

import "github.com/konstantinlevin77/courseManagement/internal/repository"

var App *AppConfig

type AppConfig struct {
	DB repository.DatabaseRepo
}

func NewApp(repo repository.DatabaseRepo) {

	App = &AppConfig{}
	App.DB = repo

}
