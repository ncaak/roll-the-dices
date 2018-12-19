package main

import (
	"net/http"
	"log"
	"github.com/ncaak/roll-the-dices/lib"
)

func tokenListener(w http.ResponseWriter, r *http.Request) {
	log.Print("request received: ", r)
	w.Write([]byte("it's alive!\n"))
}

func main() {

	server.Listen("test", "8443")
	server.Run(tokenListener)
   
}