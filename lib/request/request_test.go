package request

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"testing"
)

// Tests basic request to get last updates
func TestRequestGetUpdates(t *testing.T) {
	var test = "ENV_DEV"
	var settings = config.GetSettings(test)
	t.Logf("Test request configuration for environment: %s", test)
	var req = Init(settings.Api)
	// Pings the opened database to check if everything was fine
	var result = req.GetUpdates(1)
	t.Logf("Result: Updates returned: %+v", result)
}
