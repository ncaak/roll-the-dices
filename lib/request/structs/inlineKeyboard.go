package structs

import (
	"encoding/json"
)

// Basic set of dice to be displayed as available buttons
var baseDice = [6]string{"1d20", "1d12", "1d10", "1d8", "1d6", "1d4"}

// Inline Keyboard structures are used in Reply structures to handle user feedback
//
// InlineKeyBoard
//	\--> []Rows
//		\--> []Buttons
//			\--> Text
//			\--> Callback
//

type button struct {
	Text     string `json:"text"`
	Callback string `json:"callback_data"`
}

type row []button

type inlineKeyboard struct {
	Rows []row `json:"inline_keyboard"`
}

// --- Exported functions for the structure ---

func NewDiceKeyboard() string {
	js, err := json.Marshal(genDefaultKeyboard())
	if err != nil {
		panic(err.Error())
	}

	return string(js)
}

// --- Auxiliar functions ---

func genDefaultButtons() (rw row) {
	for _, die := range baseDice {
		rw = append(rw, button{die, "callback"})
	}
	return
}

func genDefaultKeyboard() inlineKeyboard {
	return inlineKeyboard{[]row{genDefaultButtons()}}
}
