package request

import (
	"bytes"
	"encoding/json"
	"github.com/ncaak/roll-the-dices/lib/request/structs"
	"net/http"
)

const KBD_MSG = "Escoge un dado para lanzar"
const MARKDOWN = "MarkDown"

// --- Responses ---

// Returns a reply message with an inline keyboard
func keyboardReply(r structs.Result) *bytes.Buffer {
	return encodeReply(structs.Reply{r.GetChatId(), KBD_MSG, r.GetReplyId(), "", structs.NewDiceKeyboard()})
}

// Returns a markdown format reply message
func markdownReply(r structs.Result, text string) *bytes.Buffer {
	return encodeReply(structs.Reply{r.GetChatId(), text, r.GetReplyId(), MARKDOWN, ""})
}

// Returns a plain text reply message
func textReply(r structs.Result, text string) *bytes.Buffer {
	return encodeReply(structs.Reply{r.GetChatId(), text, r.GetReplyId(), "", ""})
}

// Auxiliar function to encode a structure and returning a buffer suited to be sent in a request
func encodeReply(reply structs.Reply) *bytes.Buffer {
	jsonReply, err := json.Marshal(reply)
	if err != nil {
		panic(err.Error())
	}

	return bytes.NewBuffer(jsonReply)
}

// --- Requests ---

// Parse the response of GetUpdates request and model the data using Update structure
func getParsedUpdates(response *http.Response) []structs.Result {
	var updates = structs.Update{}
	json.NewDecoder(response.Body).Decode(&updates)
	return updates.Result
}
