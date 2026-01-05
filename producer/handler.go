package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	var event LocationEvent
	json.NewDecoder(r.Body).Decode(&event)
	event.Timestamp = time.Now().Unix()

	err := Publish(event)
	if err != nil {
		http.Error(w, "Kafka error", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
