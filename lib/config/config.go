package config

import (
	"log"
	"os"
)

// Configuration structure to handle API properties
type API struct {
	Url string
}

// Configuration structure to handle database properties
type DB struct {
	Credentials string
	Access      string
}

// Configuration structure to be handled by other modules
type Config struct {
	Api      API
	DataBase DB
}

// Retrieves settings from environment variable and builds config structure
func GetSettings() Config {
	var cfg = Config{}

	cfg.Api.Url = os.Getenv("API_URL")
	cfg.DataBase.Credentials = os.Getenv("DATABASE_CREDENTIALS")
	cfg.DataBase.Access = os.Getenv("DATABASE_ACCESS")

	if cfg.Api.Url == "" || cfg.DataBase.Credentials == "" || cfg.DataBase.Access == "" {
		log.Println("[ERR] Retrieving configuration failed")
	}

	return cfg
}
