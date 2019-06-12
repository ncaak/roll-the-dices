package config

import (
	"encoding/json"
	"os"
)

const cfgPath = "config/"
const requestsToken = "token"
const dbCredentials = "dbKey"

// Configuration structure to be handled by other modules
type config struct {
	Token string `json:"token"`
	Dbkey string `json:"dbKey"`
}

// Retrieves settings from environment variable and builds config structure
func GetSettings(env string) config {
	var cfg = config{}
	var settings = []byte(os.Getenv(env))
	
	if err := json.Unmarshal(settings, &cfg); err != nil {
		panic(err)
	}
	return cfg
}
