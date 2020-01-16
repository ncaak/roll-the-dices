package request

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"testing"
)

/*
 * Mocks
 */
type mockSettings struct{}

func (s mockSettings) GetApiUrl() string { return "https://localhost" }

/*
 * Tests
 */
func initApi() core {
	cfg, _ := config.GetSettings()
	return Init(cfg)
}

func panicRecover(t *testing.T) func() {
	return func() {
		if r := recover(); r == nil {
			t.Error("ERROR :: Wrong query to mock core")
		}
	}
}

func TestGetUpdatesOK(t *testing.T) {
	var api = initApi()
	api.GetUpdates(0)
}

func TestGetUpdatesKO(t *testing.T) {
	defer panicRecover(t)()
	var api = Init(mockSettings{})
	api.GetUpdates(0)
}
