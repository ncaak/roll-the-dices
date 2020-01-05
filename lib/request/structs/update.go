package structs

import (
	"encoding/json"
	"io"
)

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
	Message  Message  `json:"message,omitempty"`
	Callback Callback `json:"callback_query,omitempty"`
	UpdateId int      `json:"update_id"`
}

type Update struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

// --- Exported methods for the structure ---

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

// Parse the response of GetUpdates request and model the data using Update structure
func DecodeUpdates(str io.ReadCloser) []Result {
	var updates = Update{}
	if err := json.NewDecoder(str).Decode(&updates); err != nil {
		panic(err.Error())
	}
	return updates.Result
}
