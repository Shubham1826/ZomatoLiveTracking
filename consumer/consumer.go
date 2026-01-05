package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func Consume() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "delivery.location.v1",
		GroupID: "location-consumers",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			continue
		}

		var event LocationEvent
		json.Unmarshal(m.Value, &event)

		key := "ws:order:" + event.OrderID
		connIDs, _ := rdb.SMembers(context.Background(), key).Result()

		for _, connID := range connIDs {
			mu.RLock()
			conn := connections[connID]
			mu.RUnlock()

			if conn == nil {
				continue
			}

			if err := conn.WriteJSON(event); err != nil {
				conn.Close()
				mu.Lock()
				delete(connections, connID)
				mu.Unlock()

				rdb.SRem(context.Background(), key, connID)
			}
		}
	}
}

