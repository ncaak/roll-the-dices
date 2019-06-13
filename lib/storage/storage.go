package storage

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const tbUpdates = "Updates"

// Structure to handle operations with database
type dataBase struct {
	core *sql.DB
}

// Initialize database opening it and waiting for next operations to be done
func Init(cfg config.Config) dataBase {
	fmt.Println("New connection to database")

	db, err := sql.Open("mysql", cfg.Dbkey + "@/pifiabot")
	if err != nil {
		panic(err.Error())
	}
	
	fmt.Println("Connection to database successful")
	return dataBase{db}
}

// Close database operations
func (db *dataBase) Close () {
	defer db.core.Close()
	fmt.Println("Connection to database closed")
}

// Retrieves offset value saved into database previously
func (db *dataBase) GetOffset () int {
	var results = db.query(fmt.Sprintf("SELECT * FROM %s", tbUpdates))

	var offset int
	for results.Next() {
		results.Scan(&offset)
	}
	
	return offset
}

// Saves offset value into database
func (db *dataBase) SetOffset (offset string) {
	db.query(fmt.Sprintf("UPDATE %s SET offset=%s", tbUpdates, offset))
}

// Non exported function to send the queries to database
func (db *dataBase) query (queryString string) *sql.Rows {
	rows, err := db.core.Query(queryString)
	if err != nil {
		panic(err.Error())
	}

	return rows
}

