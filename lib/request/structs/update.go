package structs

const COMMAND_TYPE = "bot_command"

// Results structures retrieved from GetUpdates request
// Exported methods included to easily access required information
//
// Update
//	\--> Result
//		\--> Message
//			\--> Chat
//			\--> Entities
//			\--> From
//
type Chat struct {
	FirstName string `json:"first_name"`
	Id        int    `json:"id"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Username  string `json:"username"`
}

type Entities struct {
	Length int    `json:"length"`
	Offset int    `json:"offset"`
	Type   string `json:"type"`
}

type From struct {
	FirstName    string `json:"first_name"`
	Id           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	LanguageCode string `json:"language_code"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
}

type Message struct {
	Chat      Chat       `json:"chat"`
	Date      int        `json:"date"`
	Entities  []Entities `json:"entities"`
	From      From       `json:"from"`
	MessageId int        `json:"message_id"`
	Text      string     `json:"text"`
}

type Result struct {
	Message  Message `json:"message"`
	UpdateId int     `json:"update_id"`
}

type Update struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

// --- Exported methods for the structure ---

// Returning source Chat identifier to send the reply
func (result *Result) GetChatId() int {
	return result.Message.Chat.Id
}

// Retrieves the message content filtered as it is supposed to be a command
func (result *Result) GetCommand() string {
	return result.Message.Text
}

// Returning source Message identifier to send the reply to the original sender
func (result *Result) GetReplyId() int {
	return result.Message.MessageId
}

// Returns if a message is a command or only another message type which will be ignored
func (result *Result) IsCommand() (command bool) {
	if ent := result.Message.Entities; len(ent) > 0 && ent[0].Type == COMMAND_TYPE {
		command = true
	}
	return
}
