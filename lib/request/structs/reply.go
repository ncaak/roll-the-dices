package structs

import (
	"bytes"
	"encoding/json"
)

const MARKDOWN = "MarkDown"
const KBD_MSG = "Escoge un dado para lanzar"

// Reply structures are used in SendMessage requests
type Reply struct {
	ChatId  int    `json:"chat_id"`
	Text    string `json:"text"`
	ReplyId int    `json:"reply_to_message_id"`
	Parse   string `json:"parse_mode"`
	Markup  string `json:"reply_markup"`
}

/*
 * Structure initializers
 * Returns encoded structures to use within request library
 */

// Main structure initializer with basic data
func initReply(chatId int, replyId int) (r Reply) {
	r.ChatId = chatId
	r.ReplyId = replyId
	return
}

// Returns a reply with a simple unformatted text
func InitBasicReply(chatId int, replyId int, text string) *bytes.Buffer {
	var r = initReply(chatId, replyId)
	r.Text = text
	return r.encode()
}

// Returns a reply with an inline keyboard
func InitKeyboard(chatId int, replyId int) *bytes.Buffer {
	var r = initReply(chatId, replyId)
	r.Text = KBD_MSG
	r.Markup = NewDiceKeyboard()
	return r.encode()
}

// Returns a reply with markdown formatted text
func InitMarkdown(chatId int, replyId int, text string) *bytes.Buffer {
	var r = initReply(chatId, replyId)
	r.Text = text
	r.Parse = MARKDOWN
	return r.encode()
}

/*
 * Structure methods
 */

// Auxiliar function to encode a structure and returning a buffer suited to be sent in a request
func (r Reply) encode() *bytes.Buffer {
	jsonReply, err := json.Marshal(r)
	if err != nil {
		panic(err.Error())
	}

	return bytes.NewBuffer(jsonReply)
}
