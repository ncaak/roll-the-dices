package config

import (
	"testing"
)

// Tests retrieving API configuration from DEV environment variables
func TestConfigAPI(t *testing.T) {
	var test = "ENV_DEV"
	t.Logf("Test configuration for environment: %s", test)
	t.Log("Expected result: 'API: {BaseUrl:<baseUrl> Token:<token>}'")

	var result = GetSettings(test)
	t.Logf("Result: API: %+v", result.Api)
}

// Tests retrieving Database configuration from DEV environment variables
func TestConfigDB(t *testing.T) {
	var test = "ENV_DEV"
	t.Logf("Test configuration for environment: %s", test)
	t.Log("Expected result: 'Database: {Type:<type> User:<user> Pass:<pass> Name:<name> OffsetTable:<table> OffsetColumn:<column>}'")

	var result = GetSettings(test)
	t.Logf("Result: DataBase: %+v", result.DataBase)
}
