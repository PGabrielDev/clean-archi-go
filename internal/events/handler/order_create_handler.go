package handler

import (
	"encoding/json"
	"github.com/PGabrielDev/clean-archi-go/pkg/events"
	"github.com/streadway/amqp"
	"sync"
)

type RabbitMQHandler struct {
	RabbitMQChannel *amqp.Channel
}

func (r *RabbitMQHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	body, _ := json.Marshal(event.GetPayload())
	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	r.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
