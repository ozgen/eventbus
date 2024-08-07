package eventbus

import (
	"log"
	"os"
	"sync"
)

// Event defines the structure for all events.
type Event struct {
	Type    string
	Payload interface{}
}

// Listener defines the method signature for event listeners.
type Listener func(Event)

// EventBus manages listeners and dispatches events.
type EventBus struct {
	listeners map[string][]Listener
	lock      sync.Mutex
	wg        sync.WaitGroup
	logger    *log.Logger
}

// New creates a new EventBus with default logging.
func New() *EventBus {
	logger := log.New(os.Stdout, "EventBus: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &EventBus{
		listeners: make(map[string][]Listener),
		logger:    logger,
	}
}

// Subscribe adds a listener for a given event type.
func (eb *EventBus) Subscribe(eventType string, listener Listener) {
	eb.lock.Lock()
	defer eb.lock.Unlock()
	wrappedListener := func(e Event) {
		eb.wg.Add(1)
		go func() {
			defer eb.wg.Done()
			listener(e)
		}()
	}
	eb.listeners[eventType] = append(eb.listeners[eventType], wrappedListener)
	eb.logger.Printf("Subscribed new listener to event type '%s'", eventType)
}

// Publish sends an event to all registered listeners for the event's type.
func (eb *EventBus) Publish(event Event) {
	eb.lock.Lock()
	listeners, ok := eb.listeners[event.Type]
	eb.lock.Unlock()

	if ok {
		eb.logger.Printf("Publishing event of type '%s' to %d listeners", event.Type, len(listeners))
		for _, listener := range listeners {
			listener(event) // Listener already in a goroutine via subscribe
		}
	} else {
		eb.logger.Printf("No listeners for event type '%s'", event.Type)
	}
}

// SetLogger allows external configuration of the logger.
func (eb *EventBus) SetLogger(logger *log.Logger) {
	eb.lock.Lock()
	defer eb.lock.Unlock()
	eb.logger = logger
}

// Wait blocks until all events have been processed.
func (eb *EventBus) Wait() {
	eb.wg.Wait()
}
