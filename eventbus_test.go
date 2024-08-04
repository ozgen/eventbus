package eventbus

import (
	"sync"
	"testing"
	"time"
)

// TestEventBus tests single listener functionality.
func TestEventBus(t *testing.T) {
	eb := New()
	received := make(chan bool, 1) // Buffered channel to prevent goroutine from hanging

	eb.Subscribe("test", func(e Event) {
		if s, ok := e.Payload.(string); !ok || s != "hello" {
			t.Errorf("Expected 'hello', got '%v'", e.Payload)
		}
		received <- true
	})

	eb.Publish(Event{"test", "hello"})
	select {
	case <-received:
	case <-time.After(time.Second):
		t.Fatal("Did not receive event in time")
	}
}

// TestEventBusMultipleListeners tests the delivery of events to multiple listeners.
func TestEventBusMultipleListeners(t *testing.T) {
	eb := New()
	var wg sync.WaitGroup
	received1 := make(chan bool, 1)
	received2 := make(chan bool, 1)

	wg.Add(2) // Expect two listeners

	eb.Subscribe("test", func(e Event) {
		defer wg.Done()
		if s, ok := e.Payload.(string); !ok || s != "hello" {
			t.Errorf("Listener 1: Expected 'hello', got '%v'", e.Payload)
		}
		received1 <- true
	})

	eb.Subscribe("test", func(e Event) {
		defer wg.Done()
		if s, ok := e.Payload.(string); !ok || s != "hello" {
			t.Errorf("Listener 2: Expected 'hello', got '%v'", e.Payload)
		}
		received2 <- true
	})

	eb.Publish(Event{"test", "hello"})
	wg.Wait() // Wait for all listeners to complete

	// Check if both listeners received the message
	select {
	case <-received1:
	default:
		t.Fatal("Listener 1 did not receive event in time")
	}

	select {
	case <-received2:
	default:
		t.Fatal("Listener 2 did not receive event in time")
	}
}

// Additional test to ensure graceful shutdown of EventBus
func TestEventBusShutdown(t *testing.T) {
	eb := New()
	done := make(chan bool, 1)
	count := 0

	eb.Subscribe("shutdown", func(e Event) {
		count++
		if count == 10 {
			done <- true
		}
	})

	for i := 0; i < 10; i++ {
		eb.Publish(Event{"shutdown", i})
	}

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Did not finish processing all events before shutdown")
	}
}
