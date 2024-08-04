package eventbus

import (
	"testing"
	"time"
)

func TestEventBus(t *testing.T) {
	eb := New()
	received := make(chan bool)

	eb.Subscribe("test", func(e Event) {
		if e.Payload.(string) != "hello" {
			t.Errorf("Expected 'hello', got '%s'", e.Payload)
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

func TestEventBusMultipleListeners(t *testing.T) {
	eb := New()
	received1 := make(chan bool)
	received2 := make(chan bool)

	eb.Subscribe("test", func(e Event) {
		if e.Payload.(string) != "hello" {
			t.Errorf("Listener 1: Expected 'hello', got '%s'", e.Payload)
		}
		received1 <- true
	})

	eb.Subscribe("test", func(e Event) {
		if e.Payload.(string) != "hello" {
			t.Errorf("Listener 2: Expected 'hello', got '%s'", e.Payload)
		}
		received2 <- true
	})

	eb.Publish(Event{"test", "hello"})
	select {
	case <-received1:
	case <-time.After(time.Second):
		t.Fatal("Listener 1 did not receive event in time")
	}

	select {
	case <-received2:
	case <-time.After(time.Second):
		t.Fatal("Listener 2 did not receive event in time")
	}
}
