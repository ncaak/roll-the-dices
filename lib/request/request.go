package request

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
)

// --- GET ---

// Sends a GET request to server to retrieve updates from the Offset
// offset - Last retrieved update message Id
func (api api) GetUpdates(offset int) []structs.Result {
	var query = map[string]string{
		"offset": fmt.Sprintf("%d", offset),
	}

	resp := api.get("getUpdates", buildQuerystr(query))

	defer resp.Body.Close()
	return structs.DecodeUpdates(resp.Body)
}

// --- POST ---

func (api api) EditReplyKeyboard(msg structs.Result) {
	//
	//		"POST",
	//		fmt.Sprintf("%s%s", api.url, "editMessageReplyMarkup"),
}

// Sends a POST request to server to deliver the message with markdown style
// message - Structure with the required fields to send the reply like Chat and Message indentifier
// text - Messsae plain text to be sent as part of the reply
func (api api) ReplyHelp(msg structs.Result, text string) {
	api.post("sendMessage", structs.InitMarkdownReply(msg, text))
}

// Sends a POST request to server to deliver the message with an inline keyboard
// message - Structure with the required fields to send the reply like Chat and Message indentifier
func (api *api) ReplyInlineKeyboard(msg structs.Result) {
	api.post("sendMessage", structs.InitKeyboardReply(msg))
}

// Sends a POST request to server to deliver the reply to a message
// message - Structure with the required fields to send the reply like Chat and Message indentifier
// text - Messsae plain text to be sent as part of the reply
func (api *api) Reply(msg structs.Result, text string) {
	api.post("sendMessage", structs.InitReply(msg, text))
}
