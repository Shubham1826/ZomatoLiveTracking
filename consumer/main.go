package main

import (
	"log"
	"net/http"
)

func main() {
	go Consume()
	http.HandleFunc("/ws", WSHandler)
	http.ListenAndServe(":8090", nil)
	log.Println("consumer started on : 8080")
}
