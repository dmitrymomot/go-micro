package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				nc.Publish("current_time", []byte(fmt.Sprintf("Current time is %s", time.Now().String())))
			}
		}
	}()

	nc.Subscribe("current_time", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	wg.Wait()
}
