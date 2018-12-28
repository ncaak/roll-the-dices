package main

import (
	"github.com/ncaak/roll-the-dices/lib/config"
	"github.com/ncaak/roll-the-dices/lib/storage"
	"github.com/ncaak/roll-the-dices/structs/update"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

func getUpdates(offset int) []update.Result {
//	var url = fmt.Sprintf("https://api.telegram.org/bot%s/%s", config.GetToken(), "getUpdates")
	
	resp, err := http.Get(fmt.Sprintf(
		"https://api.telegram.org/bot%s/getUpdates?offset=%d",
		config.GetToken(),
		offset + 1,
	))
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	
	response := update.Update{}
	json.NewDecoder(resp.Body).Decode(&response)

	// debug response print
//	marshal, _ := json.Marshal(response)
//	fmt.Print(string(marshal))

	return response.Result
}

func main() {
	log.Println("beginning routine")

	var offset = storage.GetUpdateOffset()

	log.Print(getUpdates(offset))

	storage.Close()

	log.Println("ending routine")
}
