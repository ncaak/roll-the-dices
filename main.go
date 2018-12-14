package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

func tokenListener(w http.ResponseWriter, r *http.Request) {
	log.Print("request received: ", r)
	w.Write([]byte("it's alive!\n"))
}

func main() {
	// Recover certificate and private key for HTTPS server
	data, err := ioutil.ReadFile("certs/cert.pem")
	if err != nil {
		log.Fatal("no data", err)
	}

	log.Print(string(data))

	// Creates a simple web server
	http.HandleFunc("/token", tokenListener)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("oooops: ", err)
	}
}