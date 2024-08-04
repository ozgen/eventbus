package main

import (
	"fmt"
	"github.com/ozgen/eventbus"
)

func main() {
	// Create a new instance of EventBus
	bus := eventbus.New()

	// Subscribe to an event type with a simple handler
	bus.Subscribe("greeting", func(e eventbus.Event) {
		fmt.Println("Received greeting:", e.Payload)
	})

	// Publish an event
	bus.Publish(eventbus.Event{Type: "greeting", Payload: "Hello, World!"})
}
