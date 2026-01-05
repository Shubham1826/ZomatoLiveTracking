package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var (
	connections = make(map[string]*websocket.Conn) // connID → conn
	mu          sync.RWMutex
)

func WSHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	connID := uuid.NewString()

	// Store in memory
	mu.Lock()
	connections[connID] = conn
	mu.Unlock()

	// Store in Redis
	key := "ws:order:" + orderID
	rdb.SAdd(context.Background(), key, connID)

	log.Println("WS connected", orderID, connID)

	// Cleanup on disconnect
	go func() {
		defer conn.Close()
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				mu.Lock()
				delete(connections, connID)
				mu.Unlock()

				rdb.SRem(context.Background(), key, connID)
				log.Println("WS disconnected", connID)
				break
			}
		}
	}()
}

