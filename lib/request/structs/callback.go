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
	ReplyMsg  message        `json:"reply_to_message"`
	Text      string         `json:"text"`
}

type callback struct {
	ChatInst string    `json:"chat_instance"`
	Data     string    `json:"data"`
	From     from      `json:"from"`
	Id       string    `json:"id"`
	Message  cbMessage `json:"message"`
}
