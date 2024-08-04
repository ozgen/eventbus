package main

import (
	"fmt"
	"github.com/ozgen/eventbus"
	"time"
)

// SimpleEvent defines a basic structure for an event's data.
type SimpleEvent struct {
	Payload interface{}
}

// TestEventListener is a listener that handles test events.
type TestEventListener struct{}

// Handle processes events directed to this listener.
func (l *TestEventListener) Handle(event eventbus.Event) {
	if se, ok := event.Payload.(*SimpleEvent); ok {
		fmt.Printf("Handled test event with data: %v\n", se.Payload)
	} else {
		fmt.Println("Failed to handle test event: incorrect type")
	}
}

func main() {
	// Create a new instance of EventBus
	bus := eventbus.New()

	// Create and register a test event listener
	testListener := &TestEventListener{}
	bus.Subscribe("testEvent", testListener.Handle)

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

	// Publish a test event
	bus.Publish(eventbus.Event{Type: "testEvent", Payload: &SimpleEvent{Payload: "Test event payload"}})

	// Wait a bit to ensure the event is processed before the program exits
	time.Sleep(1 * time.Second)

	// Use the WaitGroup to ensure all handlers complete before the program exits
	// This is especially useful in larger applications where you need to ensure all processing is complete
	bus.Wait()
}
