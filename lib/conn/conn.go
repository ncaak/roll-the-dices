package conn

import(
	"github.com/ncaak/roll-the-dices/structs/update"
	"github.com/ncaak/roll-the-dices/lib/config"
	"net/http"
	"time"
	"encoding/json"
	"bytes"
	"fmt"
)

func send(r *http.Request) *http.Response {
	var client = &http.Client{}

	client.Timeout = 30 * time.Second

	resp, err := client.Do(r)
	if err != nil {
		panic(err.Error())
	}

	return resp
}

func GetUpdates(env string, offset int) []update.Result {
	var url = fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", config.GetSettings(env).Token, offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	var resp = send(req)

	defer resp.Body.Close()
	
	response := update.Update{}
	json.NewDecoder(resp.Body).Decode(&response) 

	return response.Result
}

func SendReply(env string, chatId int, msgText string, replyId int) {
	var url = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.GetSettings(env).Token)

	type Message struct {
		ChatId int `json:"chat_id"`
		Text string `json:"text"`
		ReplyId int `json:"reply_to_message_id"`
	}


	msg := Message{chatId, msgText, replyId}
	jsonString, _ := json.Marshal(msg)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	if err != nil {
		panic(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	var resp = send(req)
	
	defer resp.Body.Close()
}
