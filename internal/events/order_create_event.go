package events

import "time"

type OrderCreatedEvent struct {
	Name    string
	Payload interface{}
}

func (o *OrderCreatedEvent) GetName() string {
	return o.Name
}
func (o *OrderCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}
func (o *OrderCreatedEvent) GetPayload() interface{} {
	return o.Payload
}
func (o *OrderCreatedEvent) SetPayload(payload interface{}) {
	o.Payload = payload
}


