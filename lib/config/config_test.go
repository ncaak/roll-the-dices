package config

import (
	"testing"
)

// Tests basic configuration retrieve from DEV environment
func TestConfigBasic(t *testing.T) {
	var test = "ENV_DEV"
	t.Logf("Test configuration for environment: %s", test)
	t.Log("Expected result: 'Token: <token>, DbKey: <dbkey>'")
	
	var result = setEnvironment(test)
	t.Logf("Result: Token: %s, DbKey: %s", result.Token, result.Dbkey)
}

