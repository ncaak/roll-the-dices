package config

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration structure to handle API properties
type API struct {
	BaseUrl string `json:"base_url"`
	Token   string `json:"token"`
}

// Configuration structure to handle database properties
type DB struct {
	Type         string `json:"type"`
	Credentials	 string `json:"cred"`
	Access		 string `json:"access"`
}

// Configuration structure to be handled by other modules
type Config struct {
	Api      API `json:"api"`
	DataBase DB  `json:"database"`
}

// Retrieves settings from environment variable and builds config structure
func GetSettings(env string) Config {
	var cfg = Config{}
	var settings = []byte(os.Getenv(env))

	if err := json.Unmarshal(settings, &cfg); err != nil {
		log.Println("[ERR] Retrieving configuration failed")
		panic(err)
	}
	return cfg
}
