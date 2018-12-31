package conn

import(
	"github.com/ncaak/roll-the-dices/structs/update"
	"github.com/ncaak/roll-the-dices/lib/config"
	"net/http"
	"encoding/json"
	"fmt"
)

func send(r *http.Request) *http.Response {
	var client = &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		panic(err.Error())
	}

	return resp
}

func GetUpdates(offset int) []update.Result {
	var url = fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", config.GetToken(), offset)

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
