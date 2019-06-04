package config

import (
	"io/ioutil"
	"strings"
)

const cfgPath = "config/"
const requestsToken = "token"
const dbCredentials = "dbKey"

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
