package main

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

var wg sync.WaitGroup

func main() {
	nc, err := nats.Connect("nats://nats-cluster:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	wg.Add(1)
	s, err := nc.Subscribe("current_time", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}
	defer s.Unsubscribe()

	wg.Wait()
}
