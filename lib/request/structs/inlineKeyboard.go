package structs

import (
	"encoding/json"
)

// Inline Keyboard structures are used in Reply structures to handle user feedback
//
// InlineKeyBoard
//	\--> []Rows
//		\--> []Buttons
//			\--> Text
//			\--> Callback
//

type key struct {
	Text     string `json:"text"`
	Callback string `json:"callback_data"`
}

type row []key

type inlineKeyboard struct {
	Rows []row `json:"inline_keyboard"`
}

/*
 * Encodad keyboard information to be included in replies structures
 */

// Returns a complete inline keyboard with the dice keys on rows of MAX_KEYS_ON_A_ROW
func NewDiceKeyboard() string {
	const MAX_KEYS_ON_A_ROW = 4
	var keyboardRows []row

	packKeysOnRowsOf(MAX_KEYS_ON_A_ROW, getDiceKeys(), &keyboardRows)
	
	js, _ := json.Marshal(inlineKeyboard{keyboardRows})
	return string(js)
}

func NewKeyboard(buttons map[string]string) string {
	var r = row{}

	for k, v := range buttons {
		r = append(r, key{k, v})
	}

	js, _ := json.Marshal(inlineKeyboard{[]row{r}})
	return string(js)
}

/*
 * Auxiliar functions to generate keyboard structures
 */

// Generate keys for dice options
func getDiceKeys() (r row) {
	var options = [...]string{"1d100", "1d20", "1d12", "1d10", "1d8", "1d6", "1d4", "1d3"}
	
	for _, die := range options {
		r = append(r, key{die, "/tira " + die})
	}
	return
}

// It packs the keys on an arbitrary number of keys per row
func packKeysOnRowsOf(rowMax int, keys []key, keyboard *[]row) {
	if rowMax > len(keys) {
		rowMax = len(keys)
	}
	
	*keyboard = append(*keyboard, keys[:rowMax])

	if len(keys[rowMax:]) > 0 {
		packKeysOnRowsOf(rowMax, keys[rowMax:], keyboard)
	}
}
