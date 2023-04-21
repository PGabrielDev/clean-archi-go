package events

import (
	"time"
)

type OrderCreatedEvent struct {
	Name     string
	DateTime time.Time
	Payload  interface{}
}

func (e OrderCreatedEvent) GetName() string {
	return e.Name
}

func (e OrderCreatedEvent) GetDateTime() time.Time {
	return e.DateTime
}

func (e OrderCreatedEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreatedEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}
