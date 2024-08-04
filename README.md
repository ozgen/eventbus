Certainly! To update your README for the `EventBus` project, we should include the new `Wait()` method that ensures all events are processed before an application exits. This can be critical for understanding and using the library effectively, especially in applications that need to ensure all events are handled properly before termination or during a controlled shutdown.

### Updated README.md for EventBus

Here is an updated version of the `README.md`:

```markdown
# EventBus

EventBus is a lightweight, efficient, and easy-to-use publish/subscribe event bus system implemented in Go. It allows components within an application to communicate with each other using events without knowing the details of who is handling these events.

## Installation

To use EventBus in your Go projects, simply install it using `go get`:

```bash
go get github.com/ozgen/eventbus@v0.1.2
```

## Usage

Here’s a simple example of how to use the EventBus:

```go
package main

import (
    "github.com/ozgen/eventbus"
    "fmt"
)

func main() {
    bus := eventbus.New()

    bus.Subscribe("myEvent", func(e eventbus.Event) {
        fmt.Printf("Event Received: %v\n", e.Payload)
    })

    bus.Publish(eventbus.Event{Type: "myEvent", Payload: "Hello, World!"})
    
    // Wait for all event handlers to complete before exiting
    bus.Wait()
}
```

This will output:

```
Event Received: Hello, World!
```

## API Documentation

- `New() *EventBus`: Initializes and returns a new instance of EventBus.
- `Subscribe(eventType string, listener Listener)`: Registers a listener that gets called when an event of the specified type is published.
- `Publish(event Event)`: Publishes an event to all registered listeners of the event's type.
- `Wait()`: Blocks until all handlers have completed processing. This is useful for graceful shutdowns and ensuring all events are processed before an application exits.

## Contributing

Contributions are welcome! Please feel free to submit pull requests, create issues, or provide feedback.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details.
