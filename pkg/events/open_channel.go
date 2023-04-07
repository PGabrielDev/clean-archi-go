package events

import "github.com/streadway/amqp"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}
