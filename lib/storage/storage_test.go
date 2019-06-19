package storage

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"testing"
)

// Tests basic open and close database
func TestStorageBasic(t *testing.T) {
	var test = "ENV_DEV"
	var settings = config.GetSettings(test)
	t.Logf("Test database configuration for environment: %s", test)
	var db = Init(settings.DataBase)
	// Pings the opened database to check if everything was fine
	if err := db.core.Ping(); err != nil {
		t.Errorf("ERROR :: Cannot connect with database")
	}
	db.Close()
	t.Logf("Result: Conexion with the database was correct")
}

// Tests basic open, retrieve a data, and close database afterwards
func TestStorageRetrieve(t *testing.T) {
	var test = "ENV_DEV"
	var settings = config.GetSettings(test)
	t.Logf("Test database configuration for environment: %s", test)
	t.Log("Expected result: 'd1' type(d1) = int")
	var db = Init(settings.DataBase)
	// Retrieves offset value from the database
	var result = db.GetOffset()
	db.Close()
	t.Logf("Result: Offset retrieved: %d", result)
}
