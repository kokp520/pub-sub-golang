package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("START CLIENT")
	// Client()
	// fmt.Println("---- END CLIENT -----")
	
	fmt.Println(" --- START PROCESS! ----")

	pubSub := NewPubSub()

	subscribers1 := pubSub.Subscribe("new")
	subscribers2 := pubSub.Subscribe("sport")

	go func() {
		for msg := range subscribers1 {
			fmt.Printf("Subscriber 1 received message: %s\n", msg)
		}
	}()

	go func() {
		for msg := range subscribers2 {
			fmt.Printf("Subscriber 2 received message: %s\n", msg)
		}
	}()

	pubSub.Publish("new", "Latest new : go surpassed PHP in popularity!")
	pubSub.Publish("sport", "Sport update , go team wins the championship!")

	time.Sleep(3 * time.Second)
}
