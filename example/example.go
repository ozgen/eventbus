package main

import (
	"fmt"
	"github.com/ozgen/eventbus"
	"time"
)

func main() {
	// Create a new instance of EventBus
	bus := eventbus.New()

	// Subscribe to an event type with a simple handler
	bus.Subscribe("greeting", func(e eventbus.Event) {
		if str, ok := e.Payload.(string); ok {
			fmt.Println("Received greeting:", str)
		} else {
			fmt.Println("Error: Payload was not a string")
		}
	})

	// Publish an event
	bus.Publish(eventbus.Event{Type: "greeting", Payload: "Hello, World!"})

	// Wait a bit to ensure the event is processed before the program exits
	time.Sleep(1 * time.Second)

	// Use the WaitGroup to ensure all handlers complete before the program exits
	// This is especially useful in larger applications where you need to ensure all processing is complete
	bus.Wait()
}
