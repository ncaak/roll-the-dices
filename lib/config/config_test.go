package config

import (
	"testing"
)

// Tests basic configuration retrieve from DEV environment
func TestConfigBasic(t *testing.T) {
	var test = "ENV_DEV"
	t.Logf("Test configuration for environment: %s", test)
	t.Log("Expected result: 'Token: <token>, DbKey: <dbkey>'")
	
	var result = GetSettings(test)
	t.Logf("Result: Token: %s, DbKey: %s", result.Token, result.Dbkey)
	t.Logf("Result: %s", result)
}

func TestConfigDB(t *testing.T) {
	var test = "ENV_DEV"
	t.Logf("Test configuration for environment: %s", test)
	t.Log("Expected result: 'Database: {Type:<type> User:<user> Pass:<pass> Name:<name> OffsetTable:<table> OffsetColumn:<column>}'")

	var result = GetSettings(test)
	t.Logf("Result: DataBase: %+v", result.DataBase)
}
	
