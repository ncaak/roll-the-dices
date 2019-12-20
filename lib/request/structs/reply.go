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
func InitMarkdownReply(m Msg, text string) *bytes.Buffer {
	return Reply{m.GetChatId(), text, m.GetReplyId(), MARKDOWN, ""}.encode()
}

// Returns a reply message with an inline keyboard
func InitKeyboardReply(m Msg) *bytes.Buffer {
	return Reply{m.GetChatId(), KBD_MSG, m.GetReplyId(), "", NewDiceKeyboard()}.encode()
}

// Returns a plain text reply message
func InitReply(m Msg, text string) *bytes.Buffer {
	return Reply{m.GetChatId(), text, m.GetReplyId(), "", ""}.encode()
}

func TestReply(chatId int, replyId int, text string) *bytes.Buffer {
	var reply Reply
	reply.ChatId = chatId
	reply.ReplyId = replyId
	reply.Text = text
	return reply.encode()
}

// Auxiliar methods

// Auxiliar function to encode a structure and returning a buffer suited to be sent in a request
func (r Reply) encode() *bytes.Buffer {
	jsonReply, err := json.Marshal(r)
	if err != nil {
		panic(err.Error())
	}

	return bytes.NewBuffer(jsonReply)
}
