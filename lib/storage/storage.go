package storage

import (
	"log"
	"io/ioutil"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const credentialsPath = "certs/dbCredentials"

var database *sql.DB

func getFileInfo(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failure to read file", err)
	}
	return strings.TrimSuffix(string(data), "\n")
}

func connect() {
	log.Println("New connection to database")
	var userPass = getFileInfo(credentialsPath)

	db, err := sql.Open("mysql", userPass + "@/pifiabot")
	if err != nil {
		panic(err.Error())
	}
	
	log.Println("Connection to database successful")
	database = db
}

func Close() {
	defer database.Close()
	log.Println("Connection to database closed")
}

func query(queryString string) *sql.Rows {
	if database == nil {
		connect()
	}
	
	rows, err := database.Query(queryString)
	if err != nil {
		panic(err.Error())
	}

	return rows
}

func GetLastUpdateId() string {
	var results = query("SELECT * FROM telegram")

	var lastId string
	for results.Next() {
		results.Scan(&lastId)
	}

	return lastId
}

func SetLastUpdateId(updateId string) {
	query("UPDATE telegram SET offset=" + updateId)
}
