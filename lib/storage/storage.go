package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ncaak/roll-the-dices/lib/config"
)

// Structure to handle operations with database
type dataBase struct {
	core *sql.DB
	settings config.DB
}

// Initialize database opening it and saving settings retrieved from argument
func Init(cfg config.DB) dataBase {
	fmt.Println("New connection to database")

	db, err := sql.Open(cfg.Type, fmt.Sprintf("%s:%s@/%s", cfg.User, cfg.Pass, cfg.Name))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to database successful")
	return dataBase{db, cfg}
}

// Close database operations
func (db *dataBase) Close() {
	defer db.core.Close()
	fmt.Println("Connection to database closed")
}

// Retrieves offset value saved into database previously
// Uses SQL sentence: SELECT * FROM <TABLE>
// <TABLE> is retrieved from configuration set on Initialization
func (db *dataBase) GetOffset() int {
	var results = db.query(fmt.Sprintf("SELECT * FROM %s", db.settings.OffsetTable))

	var offset int
	for results.Next() {
		results.Scan(&offset)
	}

	return offset
}

// Saves offset value into database
// Uses SQL sentence: UPDATE <TABLE> SET <COLUMN>=<VALUE>
// <TABLE> and <COLUMN> are retrieved from configuration set on Initialization
// <VALUE> is retrieved from function's argument
func (db *dataBase) SetOffset(offset string) {
	db.query(fmt.Sprintf("UPDATE %s SET %s=%s", db.settings.OffsetTable, db.settings.OffsetColumn, offset))
}

// Non exported function to send the queries to database
func (db *dataBase) query(queryString string) *sql.Rows {
	rows, err := db.core.Query(queryString)
	if err != nil {
		panic(err.Error())
	}

	return rows
}
