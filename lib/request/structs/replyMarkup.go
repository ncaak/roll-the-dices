package structs

import (
	"bytes"
	"encoding/json"
)

// Reply structures are used in SendMessage requests
type ReplyMarkup struct {
	ChatId    int    `json:"chat_id"`
	MessageId int    `json:"message_id"`
	Text      string `json:"text"`
	Markup    string `json:"reply_markup"`
}

// Initializers

// Returns an edit reply message
func InitEditMarkupReply(chatId int, replyId int, text string) *bytes.Buffer {
	return ReplyMarkup{chatId, replyId, text, ""}.encode()
}
// Auxiliar methods

// Auxiliar function to encode a structure and returning a buffer suited to be sent in a request
func (r ReplyMarkup) encode() *bytes.Buffer {
	jsonReply, err := json.Marshal(r)
	if err != nil {
		panic(err.Error())
	}

	return bytes.NewBuffer(jsonReply)
}
