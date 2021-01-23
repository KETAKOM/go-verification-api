package main

import (
	"log"
	"net/http"

	"github.com/KETAKOM/echo-ddd/application/handler"
)

func main() {
	port := ":8080"
	hellowhandler := handler.NewHelloHandler()
	http.HandleFunc("/", hellowhandler.HelloHandler)

	log.Printf("listen server %v", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenServer Error: %v", err)
		return
	}
}
