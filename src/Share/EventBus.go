package share

type Event struct {
	Topic   string
	Payload interface{}
}

type EventHandler func(Event)

type EventBus interface {
	Suscribe(topic string, handler EventHandler)
	Unsuscribe(topic string, handler EventHandler)
	Publish(event Event)
}
