package config

import (
	"fmt"
	"os"
)

const API_URL = "API_URL"
const DATABASE_CREDENTIALS = "DATABASE_CREDENTIALS"
const DATABASE_ACCESS = "DATABASE_ACCESS"

/*
 * Main structure for configuration settings
 */
type Config struct {
	api struct {
		url string
	}
	db struct {
		creds  string
		access string
	}
}

func (c Config) GetUserCred() string { return c.db.creds }
func (c Config) GetDBAccess() string { return c.db.access }
func (c Config) GetApiUrl() string   { return c.api.url }

// OS Handler to get properties from envinronment variables
type osHandler struct{}

func (h osHandler) get(key string) string { return os.Getenv(key) }

/*
 * Interface to manage Handler
 */
type handler interface {
	get(string) string
}

/*
 * Initializer
 */
func GetSettings() (Config, error) {
	return newConfig(osHandler{})
}

func newConfig(h handler) (c Config, err error) {
	c.api.url = h.get(API_URL)
	if c.api.url == "" {
		return c, fmt.Errorf("There is no configuration for %s", API_URL)
	}

	c.db.creds = h.get(DATABASE_CREDENTIALS)
	if c.db.creds == "" {
		return c, fmt.Errorf("There is no configuration for %s", DATABASE_CREDENTIALS)
	}

	c.db.access = h.get(DATABASE_ACCESS)
	if c.db.access == "" {
		return c, fmt.Errorf("There is no configuration for %s", DATABASE_ACCESS)
	}

	return c, nil
}
