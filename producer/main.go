package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Initializing Kafka producer...")
	InitKafka()
	http.HandleFunc("/location", LocationHandler)
	log.Println("Producer running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Producer running on port 8080")
}
