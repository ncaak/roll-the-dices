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

// Initializers

// Returns a markdown format reply message
func InitMarkdownReply(r Result, text string) *bytes.Buffer {
	return encode(Reply{r.GetChatId(), text, r.GetReplyId(), MARKDOWN, ""})
}

// Returns a reply message with an inline keyboard
func InitKeyboardReply(r Result) *bytes.Buffer {
	return encode(Reply{r.GetChatId(), KBD_MSG, r.GetReplyId(), "", NewDiceKeyboard()})
}

// Returns a plain text reply message
func InitReply(r Result, text string) *bytes.Buffer {
	return encode(Reply{r.GetChatId(), text, r.GetReplyId(), "", ""})
}

// Auxiliar methods

// Auxiliar function to encode a structure and returning a buffer suited to be sent in a request
func encode(r Reply) *bytes.Buffer {
	jsonReply, err := json.Marshal(r)
	if err != nil {
		panic(err.Error())
	}

	return bytes.NewBuffer(jsonReply)
}
