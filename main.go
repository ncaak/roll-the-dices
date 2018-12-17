package main

import (
	"net/http"
	"log"
	//"io/ioutil"
	"github.com/ncaak/roll-the-dices/lib"
)

func tokenListener(w http.ResponseWriter, r *http.Request) {
	log.Print("request received: ", r)
	w.Write([]byte("it's alive!\n"))
}

func main() {
	// Recover certificate and private key for HTTPS server
	// data, err := ioutil.ReadFile("certs/cert.pem")
	// if err != nil {
	// 	log.Fatal("no data", err)
	// }

	server.Listen("arg", "12345")
	server.Run(tokenListener)
   
}