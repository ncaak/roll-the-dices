package request

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

/*
 * Structure to handle api requests
 */
type core struct {
	client *http.Client
	url    string
}

/*
 * Interface used to set api configuration
 */
type settings interface {
	GetApiUrl() string
}

// Initialize client to http library package and prepare package structure to be used afterwards
func Init(cfg settings) (api core) {
	api.client = &http.Client{}
	api.client.Timeout = 30 * time.Second
	api.url = cfg.GetApiUrl()
	return
}

// Base GET request handler, returns response if no error is found
func (api core) get(method string, query string) *http.Response {
	var url = fmt.Sprintf("%s%s%s", api.url, method, query)
	// Prepare the request to retrieve unreaded messages
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	// Does the request and returns the response
	resp, err := api.client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	return resp
}

// Base POST request handler, returns response if no error is found
func (api *core) post(method string, body *bytes.Buffer) *http.Response {
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

//
func (api core) sendRequest(r string) {

}

// --- Auxiliar methods ---

// Builds a query string for a GET Resquest
func buildQuerystr(values map[string]string) string {
	var query strings.Builder
	if len(values) > 0 {
		// Beginning of the query string
		query.WriteString("?")
		for k, v := range values {
			// Inserts ampersand symbol before key-value values if not the first value
			if query.Len() > 1 {
				query.WriteString("&")
			}
			fmt.Fprintf(&query, "%s=%s", k, v)
		}
	}
	return query.String()
}
