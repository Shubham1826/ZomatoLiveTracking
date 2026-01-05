package main

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Register(orderID string, connID string) {
	rdb.SAdd(context.Background(), orderID, connID)
}
