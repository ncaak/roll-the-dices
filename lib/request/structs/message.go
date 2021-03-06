package structs

// Interface for structs containing message information needed to send API requests
type Msg interface {
	GetChatId() int
	GetReplyId() int
}

// Message structures find in Callback and Result structures
//
// Message
//	\--> Chat
//	\--> Entities
//	\--> From
//
type chat struct {
	FirstName string `json:"first_name"`
	Id        int    `json:"id"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Username  string `json:"username"`
}

type entities struct {
	Length int    `json:"length"`
	Offset int    `json:"offset"`
	Type   string `json:"type"`
}

type from struct {
	FirstName    string `json:"first_name"`
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	LanguageCode string `json:"language_code"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
}

type Message struct {
	Chat      chat       `json:"chat"`
	Date      int        `json:"date"`
	Entities  []entities `json:"entities"`
	From      from       `json:"from"`
	MessageId int        `json:"message_id"`
	Text      string     `json:"text"`
}

// --- Exported methods for the structure ---

// Returning source Chat identifier to send the reply
func (msg Message) GetChatId() int {
	return msg.Chat.Id
}

func (msg Message) GetCommand() string {
	return msg.Text
}

// Returning source Message identifier to send the reply to the original sender
func (msg Message) GetReplyId() int {
	return msg.MessageId
}
