package share_infrastructure_inmemoryeventbus

import (
	"fmt"
	domain "gasto-api/src/Share"
	"sync"
)

type inMemoryEventBus struct {
	suscribers map[string][]domain.EventHandler
	lock       sync.RWMutex
}

func CreateInMemoryEventBus() domain.EventBus {
	return &inMemoryEventBus{
		suscribers: make(map[string][]domain.EventHandler),
	}
}

func (bus *inMemoryEventBus) Suscribe(topic string, handler domain.EventHandler) {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	bus.suscribers[topic] = append(bus.suscribers[topic], handler)
}

func (bus *inMemoryEventBus) Unsuscribe(topic string, handler domain.EventHandler) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	if handlers, hasHandlers := bus.suscribers[topic]; hasHandlers {
		for index, registerHandler := range handlers {
			if areSameHandler(registerHandler, handler) {
				bus.suscribers[topic] = domain.RemoveFromArrayUnsafe(handlers, index)
			}
		}
	}
}

func areSameHandler(handler1, handler2 domain.EventHandler) bool {
	return fmt.Sprintf("%v", handler1) == fmt.Sprintf("%v", handler2)
}

func (bus *inMemoryEventBus) Publish(event domain.Event) {
	bus.lock.Lock()
	defer bus.lock.RUnlock()
	if handlers, found := bus.suscribers[event.Topic]; found {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}
