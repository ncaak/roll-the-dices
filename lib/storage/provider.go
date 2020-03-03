package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const DRIVER = "mysql"

// Initialize database opening it and saving settings retrieved from argument
func Init(cfg settings) (dBase, error) {
	// Opening the database allowing to send queries
	db, err := sql.Open(DRIVER, fmt.Sprintf("%s@%s", cfg.GetUserCred(), cfg.GetDBAccess()))
	if err != nil {
		return dBase{}, err
	}
	return dBase{db, defineDB()}, nil
}

// Close database operations
func (db dBase) Close() {
	db.core.Close()
}

// Non exported function to send the queries to database
func (db dBase) query(q string) *sql.Rows {
	rows, err := db.core.Query(q)
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
