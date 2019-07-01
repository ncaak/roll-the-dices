package request

import (
	"bytes"
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/config"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
	"net/http"
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

// --- Responses ---

// Sends a GET request to server to retrieve updates from the Offset
// offset - Last retrieved update message Id
func (api *api) GetUpdates(offset int) []structs.Result {
	// Prepare the request to retrieve unreaded messages
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s%s?offset=%d", api.url, "getUpdates", offset),
		nil,
	)
	if err != nil {
		panic(err.Error())
	}
	// Handle the response and parse it to return modelled data (structs Update)
	var resp = api.send(req)
	defer resp.Body.Close()

	return getParsedUpdates(resp)
}

// --- Requests ---

// Sends a POST request to server to deliver the message with markdown style
// message - Structure with the required fields to send the reply like Chat and Message indentifier
// text - Messsae plain text to be sent as part of the reply
func (api *api) ReplyHelp(message structs.Result, text string) {
	api.sendMessage(markdownReply(message, text))
}

// Sends a POST request to server to deliver the message with an inline keyboard
// message - Structure with the required fields to send the reply like Chat and Message indentifier
func (api *api) ReplyInlineKeyboard(message structs.Result) {
	api.sendMessage(keyboardReply(message))
}

// Sends a POST request to server to deliver the reply to a message
// message - Structure with the required fields to send the reply like Chat and Message indentifier
// text - Messsae plain text to be sent as part of the reply
func (api *api) Reply(message structs.Result, text string) {
	api.sendMessage(textReply(message, text))
}

// Sends a POST request to server to deliver the message
func (api *api) sendMessage(body *bytes.Buffer) {
	// Prepare the request to send the reply to the server
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", api.url, "sendMessage"),
		body,
	)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	// No need of handle the response if there is no break in the request
	var resp = api.send(req)
	defer resp.Body.Close()
}

// Auxiliary method to send requests
func (api *api) send(r *http.Request) *http.Response {
	resp, err := api.client.Do(r)
	if err != nil {
		panic(err.Error())
	}
	return resp
}
