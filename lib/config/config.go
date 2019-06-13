package config

import (
	"encoding/json"
	"os"
)

// Configuration structure to be handled by other modules
type Config struct {
	Token string `json:"token"`
	Dbkey string `json:"dbKey"`
}

// Retrieves settings from environment variable and builds config structure
func GetSettings(env string) Config {
	var cfg = Config{}
	var settings = []byte(os.Getenv(env))

	if err := json.Unmarshal(settings, &cfg); err != nil {
		panic(err)
	}
	return cfg
}
