package bus

import (
	"sync"
)

type Event interface {
	Identifier() EventIdentifier
}

type EventHandler interface {
	Handle(event Event) error
}

type EventHandlerFunc func(event Event) error

func (f EventHandlerFunc) Handle(event Event) error {
	return f(event)
}

type EventBus struct {
	mu       sync.RWMutex
	handlers map[EventIdentifier][]EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[EventIdentifier][]EventHandler),
	}
}

func (eb *EventBus) RegisterHandler(event Event, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.handlers[event.Identifier()] = append(eb.handlers[event.Identifier()], handler)
}

func (eb *EventBus) Publish(event Event) error {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	handlers, exists := eb.handlers[event.Identifier()]
	if !exists {
		return nil
	}

	for _, handler := range handlers {
		if err := handler.Handle(event); err != nil {
			return err
		}
	}

	return nil
}
