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

func getUpdates(id string) []update.Result {
	var url = fmt.Sprintf("https://api.telegram.org/bot%s/%s", config.GetToken(), "getUpdates")
	
	resp, err := http.Get(url)
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

	var lastUpdateId = storage.GetLastUpdateId()

	log.Print(getUpdates(lastUpdateId))

	storage.Close()

	log.Println("ending routine")
}
