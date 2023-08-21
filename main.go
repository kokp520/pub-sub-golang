package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("START PROCESS!")

	pubSub := NewPubSub()

	subscribers1 := pubSub.Subscribe("new")
	subscribers2 := pubSub.Subscribe("sport")

	//add
	// var wg sync.WaitGroup
	// wg.Add(2) // Adding two goroutines

	go func() {

		// defer wg.Done() // Notify the WaitGroup that this goroutine is done

		for msg := range subscribers1 {
			fmt.Printf("Subscriber 1 received message: %s\n", msg)
		}
	}()

	go func() {

		// defer wg.Done() // Notify the WaitGroup that this goroutine is done

		for msg := range subscribers2 {
			fmt.Printf("Subscriber 2 received message: %s\n", msg)
		}
	}()

	pubSub.Publish("new", "Latest new : go surpassed PHP in popularity!")
	pubSub.Publish("sport", "Sport update , go team wins the championship!")

	// Sleep to allow goroutines to finish receiving messages
	// select {}
	// wg.Wait()
	// Close the subscribers channels after all goroutines have finished
	// close(subscribers1)
	// close(subscribers2)
}

type PubSub struct {
	mu          sync.RWMutex             // read and wrie
	subscribers map[string][]chan string // channel
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	channel := make(chan string, 10) // buffered channel t aviod blocking publishers

	ps.subscribers[topic] = append(ps.subscribers[topic], channel)
	fmt.Printf("Subscribed to topic: %s\n", topic)
	return channel
}

func (ps *PubSub) Publish(topic, message string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscribers, found := ps.subscribers[topic]
	if !found {
		fmt.Printf("No subscribers for topic: %s\n", topic)
		return
	}

	fmt.Printf("Publishing to topic: %s\n", topic)
	for _, ch := range subscribers {
		go func(c chan string, msg string) {
			c <- msg
		}(ch, message)
	}
}
