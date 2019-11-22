package storage

import (
	"fmt"
)

// Retrieves offset value saved into database previously
// Uses SQL sentence: SELECT * FROM <TABLE>
// <TABLE> is retrieved from configuration set on Initialization
func (db *dataBase) GetOffset() (offset int) {
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
func (db *dataBase) SetOffset(offset string) {
	db.query(fmt.Sprintf("UPDATE %s SET %s=%s", db.offset.table, db.offset.column, offset))
}

