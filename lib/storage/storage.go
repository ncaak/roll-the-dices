package storage

import (
	"fmt"
)

// Retrieves offset value saved into database previously
// Uses SQL sentence: SELECT * FROM <TABLE>
// <TABLE> is retrieved from configuration set on Initialization
func (db *dataBase) GetOffset() int {
	var results = db.query(fmt.Sprintf("SELECT * FROM %s", db.settings.OffsetTable))
	// Retrieve offset value from query
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

