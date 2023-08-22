package main

import (
	"fmt"
	"time"
)

func Client() {
	pubSub := NewPubSub()

	subscribers := pubSub.Subscribe("new")

	go func() {
		for msg := range subscribers {
			fmt.Printf("client received message: %s \n", msg)
		}
	}()

	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Client message %d", i)
		pubSub.Publish("new", message)
		time.Sleep(time.Second)
	}

	time.Sleep(3 * time.Second)
}
