package config

import (
	"encoding/json"
	"os"
)

type DB struct {
	Type string `json:"type"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
	OffsetTable string `json:"offset_table"`
	OffsetColumn string `json:"offset_column"`
}
// Configuration structure to be handled by other modules
type Config struct {
	Token string `json:"token"`
	Dbkey string `json:"dbKey"`
	DataBase DB	`json:"database"`
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
