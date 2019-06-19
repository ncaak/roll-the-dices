package structs

// Reply structures are used in SendMessage requests

type Reply struct {
	ChatId  int    `json:"chat_id"`
	Text    string `json:"text"`
	ReplyId int    `json:"reply_to_message_id"`
}
