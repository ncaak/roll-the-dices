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
	var endpoint = fmt.Sprintf("%s%s?offset=%d", api.url, "getUpdates", offset)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err.Error())
	}

	var resp = api.send(req)

	defer resp.Body.Close()

	response := structs.Update{}
	json.NewDecoder(resp.Body).Decode(&response)

	return response.Result
}

// Sends a POST request to server to deliver the reply to a message
// chatId - References the unique chat what the message was retreived from
// msgText - The reply in plain text to be delivered
// replyId - References the unique source message to appear like a reply
func (api *api) SendReply(chatId int, msgText string, replyId int) {
	var endpoint = fmt.Sprintf("%s%s", api.url, "sendMessage")

	type Message struct {
		ChatId  int    `json:"chat_id"`
		Text    string `json:"text"`
		ReplyId int    `json:"reply_to_message_id"`
	}

	msg := Message{chatId, msgText, replyId}
	jsonString, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonString))
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

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
