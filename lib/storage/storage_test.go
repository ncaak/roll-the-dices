package storage

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"testing"
)

/*
 * Mocks
 */
type mockSettings struct{}

func (s mockSettings) GetUserCred() string { return "" }
func (s mockSettings) GetDBAccess() string { return "" }

/*
 * Tests
 */
func initDB(t *testing.T) dBase {
	settings, _ := config.GetSettings()
	db, err := Init(settings)
	if err != nil {
		t.Errorf("ERROR :: %s", err)
	}
	return db
}

func TestInitiationOK(t *testing.T) {
	var db = initDB(t)
	defer db.Close()
	if errPing := db.core.Ping(); errPing != nil {
		t.Errorf("ERROR :: %s", errPing)
	}
}

func TestInitiationKO(t *testing.T) {
	db, _ := Init(mockSettings{})
	if err := db.core.Ping(); err == nil {
		t.Error("ERROR :: Connection with database not failed")
	}
}

func TestQueryGetOffsetOK(t *testing.T) {
	var db = initDB(t)
	defer db.Close()
	if db.GetOffset() == 0 {
		t.Error("ERROR :: Query failed and returned a 0 value")
	}
}
