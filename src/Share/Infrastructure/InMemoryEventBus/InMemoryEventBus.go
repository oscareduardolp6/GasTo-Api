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

	if handlers, found := bus.suscribers[topic]; found {
		for index, registerHandlers := range handlers {
			handlerExistsInRegisters := fmt.Sprintf("%v", registerHandlers) == fmt.Sprintf("%v", handler)
			if handlerExistsInRegisters {
				bus.suscribers[topic] = domain.RemoveFromArrayUnsafe(handlers, index)
			}
		}
	}
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
