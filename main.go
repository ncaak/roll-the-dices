package main

import (
	"github.com/ncaak/roll-the-dices/lib/storage"
	"github.com/ncaak/roll-the-dices/structs/update"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"log"
	"fmt"
)

var baseUrl = "https://api.telegram.org/bot%s/%s"

func getConfig(fileName string) string {
	data, err := ioutil.ReadFile("config/" + fileName)
	if err != nil {
		panic(err.Error())
	}
	return strings.TrimSuffix(string(data), "\n")
}

func getUpdates(id string) []update.Result {
	var url = fmt.Sprintf(baseUrl, getConfig("token"), "getUpdates")
	
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
