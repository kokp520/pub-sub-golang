package main

import (
	"fmt"
	"sync"
)

type PubSub struct {
	mu          sync.RWMutex
	subscribers map[string][]chan string
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
