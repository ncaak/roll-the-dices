package storage

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ncaak/roll-the-dices/lib/config"
)

// Structures to handle operations with database
type defaultDB struct {
	table	string
	column	string
}

type dataBase struct {
	core     *sql.DB
	offset	 defaultDB
}

// Initialize database opening it and saving settings retrieved from argument
func Init(cfg config.DB) dataBase {
	// Opening the database allowing to send queries
	db, err := sql.Open(cfg.Type, getInterface(cfg))
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

// Returns interhace string needed in database/sql module to init a database
func getInterface(cfg config.DB) string {
	return fmt.Sprintf("%s:%s@/%s", cfg.User, cfg.Pass, cfg.Name)
}

