package request

import (
	"bytes"
	"encoding/json"
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

// Sends a POST request to server to deliver the reply to a message
// chatId - References the unique chat what the message was retreived from
// msgText - The reply in plain text to be delivered
// replyId - References the unique source message to appear like a reply
func (api *api) SendReply(chatId int, msgText string, replyId int) {
	// Prepare the request to send the reply to the server
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", api.url, "sendMessage"),
		getReplyBody(chatId, msgText, replyId),
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

// Parse the response of GetUpdates request and model the data using Update structure
func getParsedUpdates(response *http.Response) []structs.Result {
	var updates = structs.Update{}
	json.NewDecoder(response.Body).Decode(&updates)
	return updates.Result
}

// Format the Reply using a struct and return the bytes object needed for the request
func getReplyBody(chatId int, msgText string, replyId int) *bytes.Buffer {
	var reply = structs.Reply{chatId, msgText, replyId}
	jsonReply, _ := json.Marshal(reply)

	return bytes.NewBuffer(jsonReply)
}
