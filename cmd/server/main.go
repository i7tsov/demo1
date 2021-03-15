package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/i7tsov/demo1/pkg/generator"
	"github.com/i7tsov/demo1/pkg/redisclient"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rc := &redisclient.RedisClient{
		Client: client,
	}
	g, err := generator.New(generator.Opts{
		Client: rc,
		Cycles: 5,
		Delay:  100 * time.Millisecond,
	})

	done := make(chan error)

	go func() { done <- g.Run() }()

	err = <-done
	if err != nil {
		log.Fatalf("Fatal error encounered: %v", err)
	}
	log.Print("Done.")
}
