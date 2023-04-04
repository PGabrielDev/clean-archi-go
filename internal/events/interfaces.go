package events

import "time"

type EventInterface interface {
	GetName() string
	GetDate() time.Time
	GetPayload() interface{}
}

type EventHandler interface {
	Handle(event EventInterface) error
}

type EventDispatcherInterface interface {
	Dispatcher(event EventInterface) error
}
