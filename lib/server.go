package server

import (
    "net/http"
    "log"
	"io/ioutil"
)

const certificatePath = "certs/cert.pem"
const privateKeyPath = "certs/private.key"

var port string
var endpoint string

type callback func (w http.ResponseWriter, r *http.Request)

func getFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failure to read file", err)
	}
	return string(data)
}

func Listen(url string, urlPort string) {
    endpoint = "/" + url
	port = ":" + urlPort
}

func Run(handler callback) {
    http.HandleFunc(endpoint, handler)
    //err := http.ListenAndServe(port, nil)
	err := http.ListenAndServeTLS(port, certificatePath, privateKeyPath, nil)
    if err != nil {
        log.Fatal("Fatal error:  ", err)
    }
}
