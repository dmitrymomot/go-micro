package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

var wg sync.WaitGroup

func main() {
	nc, err := nats.Connect("nats://nats-cluster:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	uuid := uuid.NewV1().String()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				msg := fmt.Sprintf("publisher %s: Current time is %s", uuid, time.Now().String())
				nc.Publish("current_time", []byte(msg))
			}
		}
	}()

	wg.Wait()
}
