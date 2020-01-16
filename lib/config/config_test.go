package config

import (
	"testing"
)

/*
 * Mocks
 */
type mockHandler struct{}

func (h mockHandler) get(key string) string { return "" }

/*
 * Tests
 */
func TestGetGlobalSettingsOK(t *testing.T) {
	_, err := GetSettings()
	if err != nil {
		t.Errorf("ERROR :: %s", err.Error())
	}
}

func TestGetGlobalSettingsKO(t *testing.T) {
	_, err := newConfig(mockHandler{})
	if err == nil {
		t.Error("ERROR :: Void environment variable did not trigger an error")
	}
}
