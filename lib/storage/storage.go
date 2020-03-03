package storage

import (
	"database/sql"
	"fmt"
)

/*
 * Structures to handle operations with database
 */
type defaultDB struct {
	table  string
	column string
}

type dBase struct {
	core   *sql.DB
	offset defaultDB
}

/*
 * Interface used to set database connection
 */
type settings interface {
	GetUserCred() string
	GetDBAccess() string
}

/*
 * Queries
 */

// Retrieves offset value saved into database previously
// Uses SQL sentence: SELECT * FROM <TABLE>
// <TABLE> is retrieved from configuration set on Initialization
func (db dBase) GetOffset() (offset int) {
	var results = db.query(fmt.Sprintf("SELECT * FROM %s", db.offset.table))
	// Retrieve offset value from query
	for results.Next() {
		results.Scan(&offset)
	}
	return
}

// Saves offset value into database
// Uses SQL sentence: UPDATE <TABLE> SET <COLUMN>=<VALUE>
// <TABLE> and <COLUMN> are retrieved from configuration set on Initialization
// <VALUE> is retrieved from function's argument
func (db dBase) SetOffset(offset string) {
	db.query(fmt.Sprintf("UPDATE %s SET %s=%s", db.offset.table, db.offset.column, offset))
}
