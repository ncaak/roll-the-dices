package storage

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const tbUpdates = "Updates"

var database *sql.DB

func connect() {
	fmt.Println("New connection to database")

	db, err := sql.Open("mysql", config.GetDbKey() + "@/pifiabot")
	if err != nil {
		panic(err.Error())
	}
	
	fmt.Println("Connection to database successful")
	database = db
}

func Close() {
	defer database.Close()
	fmt.Println("Connection to database closed")
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

func GetUpdateOffset() int {
	var results = query(fmt.Sprintf("SELECT * FROM %s", tbUpdates))

	var offset int
	for results.Next() {
		results.Scan(&offset)
	}
	
	return offset
}

func SetLastUpdateId(updateId string) {
	query(fmt.Sprintf("UPDATE %s SET offset=%s", tbUpdates, updateId))
}
