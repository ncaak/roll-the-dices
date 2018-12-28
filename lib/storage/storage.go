package storage

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

func connect() {
	log.Println("New connection to database")
	var userPass = config.GetDbKey()

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
