package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
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
func setEnvironment(env string) config {
	var cfg = config{}
	var settings = []byte(os.Getenv(env))
	
	if err := json.Unmarshal(settings, &cfg); err != nil {
		panic(err)
	}
	return cfg
}

func getFileInfo(fileName string) string {
	data, err := ioutil.ReadFile(cfgPath + fileName)
	if err != nil {
		panic(err.Error())
	}
	return strings.TrimSuffix(string(data), "\n")
}

func GetToken() string {
	return getFileInfo(requestsToken)
}

func GetDbKey() string {
	return getFileInfo(dbCredentials)
}
