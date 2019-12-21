package request

import (
	"fmt"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
)

/*
 * GET requests
 */

// Sends a GET request to server to retrieve updates from the Offset
// offset - Last retrieved update message Id
func (api core) GetUpdates(offset int) []structs.Result {
	var query = map[string]string{
		"offset": fmt.Sprintf("%d", offset),
	}

	resp := api.get("getUpdates", buildQuerystr(query))

	defer resp.Body.Close()
	return structs.DecodeUpdates(resp.Body)
}

/*
 * POST requests
 * Chat Id and Reply Id are necessary if the message sent has to be linked to a previous existing message
 */

// Sends a POST request to server to edit a message with a keyboard to remove it
// message - Structure with the required fields to send the reply like Chat and Message indentifier
func (api core) EditKeyboardReply(msg structs.Msg, text string) {
	api.post("editMessageText", structs.InitEditMarkupReply(msg, text))
}

// Sends a POST request to server to deliver the message with markdown style
// message - Structure with the required fields to send the reply like Chat and Message indentifier
// text - Messsae plain text to be sent as part of the reply
func (api core) ReplyMarkdown(msg structs.Msg, text string) {
	api.post("sendMessage", structs.InitMarkdownReply(msg, text))
}

// Sends a POST request to server to deliver the message with an inline keyboard
// message - Structure with the required fields to send the reply like Chat and Message indentifier
func (api core) ReplyInlineKeyboard(msg structs.Msg) {
	api.post("sendMessage", structs.InitKeyboardReply(msg))
}

// Deprecated
// Sends a POST request to server to deliver the reply to a message
// message - Structure with the required fields to send the reply like Chat and Message indentifier
// text - Messsae plain text to be sent as part of the reply
func (api core) Reply(msg structs.Msg, text string) {
	api.post("sendMessage", structs.InitReply(msg, text))
}

// Sends an unformatted basic reply
func (api core) BasicReply(chatId int, replyId int, text string) {
	api.post("sendMessage", structs.InitBasicReply(chatId, replyId, text))
}
