package request

import (
	"bytes"
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/config"
	"net/http"
	"strings"
	"time"
)

// Structure to handle operations with API
type api struct {
	client   *http.Client
	settings config.API
	url      string
}

// Initialize client to http library package and prepare package structure to be used afterwards
func Init(cfg config.API) api {
	var client = &http.Client{}
	client.Timeout = 30 * time.Second

	return api{client, cfg, fmt.Sprintf("%s%s/", cfg.BaseUrl, cfg.Token)}
}

// Base GET request handler, returns response if no error is found
func (api api) get(method string, query string) *http.Response {
	var url = fmt.Sprintf("%s%s%s", api.url, method, query)
	// Prepare the request to retrieve unreaded messages
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	// Does the request an returns the response
	resp, err := api.client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	return resp
}

// Base POST request handler, returns response if no error is found
func (api api) post(method string, body *bytes.Buffer) *http.Response {
	var url = fmt.Sprintf("%s%s", api.url, method)
	// Prepare the request to send the reply to the server
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	// Does the request an returns the response
	resp, err := api.client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	return resp
}

// --- Auxiliar methods ---

// Bulids a query string for a GET Resquest
func buildQuerystr(values map[string]string) string {
	var query strings.Builder
	if len(values) > 0 {
		// Beginning of the qdduery string
		query.WriteString("?")
		for k, v := range values {
			// Inserts ampersand symbol before key-value values if not the first value
			if query.Len() > 1 {
				query.WriteString("&")
			}
			query.WriteString(fmt.Sprintf("%s=%s", k, v))
		}
	}
	return query.String()
}
