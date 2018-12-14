package roll

import (
	"net/http"
	"log"
)

func tokenListener(w http.ResponseWriter, r *http.Request) {
	log.Print("request received: ", r)
	w.Write([]byte("it's alive!\n"))
}

func main() {
	// Creates a simple web server
	http.HandleFunc("/token", tokenListener)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("oooops: ", err)
	}
}