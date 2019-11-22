package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ncaak/roll-the-dices/lib/config"
	"log"
)

// Structures to handle operations with database
type defaultDB struct {
	table  string
	column string
}

type dataBase struct {
	core   *sql.DB
	offset defaultDB
}

// Initialize database opening it and saving settings retrieved from argument
func Init(dbconf config.DB) dataBase {
	// Opening the database allowing to send queries
	db, err := sql.Open(
		dbconf.Type,
		fmt.Sprintf("%s@%s", dbconf.Credentials, dbconf.Access),
	)
	if err != nil {
		log.Println("[ERR] Connection with database failed")
		panic(err.Error())
	}
	return dataBase{db, defineDB()}
}

// Close database operations
func (db *dataBase) Close() {
	db.core.Close()
}

// Non exported function to send the queries to database
func (db *dataBase) query(queryString string) *sql.Rows {
	rows, err := db.core.Query(queryString)
	if err != nil {
		log.Println("[ERR] Query to database failed")
		panic(err.Error())
	}
	return rows
}

// Define table and column default names depending on provider default DB installation
func defineDB() defaultDB {
	return defaultDB{"Updates", "offset"}
}
