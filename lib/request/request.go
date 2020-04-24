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

// Sends a reply to edit a previous message
func (api core) EditKeyboardReply(chatId int, replyId int, text string) {
	api.post("editMessageText", structs.InitEditMarkupReply(chatId, replyId, text))
}

// Sends an unformatted basic reply
func (api core) BasicReply(chatId int, replyId int, text string) {
	api.post("sendMessage", structs.InitBasicReply(chatId, replyId, text))
}

// Sends an inline keyboard for different basic die rolls
func (api core) DiceKeyboardReply(chatId int, replyId int) {
	api.post("sendMessage", structs.InitDiceKeyboard(chatId, replyId))
}

func (api core) KeyboardReply(chatId int, replyId int, buttons map[string]string) {
	api.post("sendMessage", structs.InitKeyboard(chatId, replyId, buttons))
}

// Sends a markdown formatted reply
func (api core) MarkdownReply(chatId int, replyId int, text string) {
	api.post("sendMessage", structs.InitMarkdown(chatId, replyId, text))
}
