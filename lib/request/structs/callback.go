package structs

// CallBack structures retrieved from GetUpdates request in a Result structure
//
// Callback
//	\--> From
//  \--> CbMessage
//		\--> Chat
//		\--> From
//		\--> InlineKeyboard
//		\--> Message
//
type cbMessage struct {
	Chat      chat           `json:"chat"`
	Date      int            `json:"date"`
	From      from           `json:"from"`
	MessageId int            `json:"message_id"`
	Markup    inlineKeyboard `json:"reply_markup"`
	ReplyMsg  Message        `json:"reply_to_message"`
	Text      string         `json:"text"`
}

type Callback struct {
	ChatInst string    `json:"chat_instance"`
	Data     string    `json:"data"`
	From     from      `json:"from"`
	Id       string    `json:"id"`
	Message  cbMessage `json:"message"`
}

// --- Exported methods for the structure ---

// Returning source Chat identifier to send the reply
func (cb Callback) GetChatId() int {
	return cb.Message.Chat.Id
}

// Returning source Message identifier to send the reply to the original sender
func (cb Callback) GetReplyId() int {
	return cb.Message.MessageId
}
