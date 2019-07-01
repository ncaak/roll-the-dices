package structs

const COMMAND_TYPE = "bot_command"

// Results structures retrieved from GetUpdates request
// Exported methods included to easily access required information
//
// Update
//	\--> Result
//		\--> Message
//		\--> Callback
//
type Result struct {
	Message  message  `json:"message,omitempty"`
	Callback callback `json:"callback_query,omitempty"`
	UpdateId int      `json:"update_id"`
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

// Returns if a message is a callback type (inline keyboard response)
func (r *Result) IsCallback() (callback bool) {
	if r.Callback.Id != "" {
		callback = true
	}
	return
}
