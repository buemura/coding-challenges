package queue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/buemura/btg-challenge/config"
	"github.com/buemura/btg-challenge/internal/modules/order"
	"github.com/buemura/btg-challenge/internal/shared"
	amqp "github.com/rabbitmq/amqp091-go"
)

const ORDER_CREATED_QUEUE = "btg.order.created"

func StartConsume() {
	// Connect to rabbitmq broker
	conn, err := amqp.Dial(config.BROKER_URL)
	shared.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open channel
	ch, err := conn.Channel()
	shared.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare/Create Queue
	_, err = ch.QueueDeclare(
		ORDER_CREATED_QUEUE, // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	shared.FailOnError(err, "Failed to declare a queue")

	// Consume Queue
	msgs, err := ch.Consume(
		ORDER_CREATED_QUEUE, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	shared.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			switch d.RoutingKey {
			case ORDER_CREATED_QUEUE:
				var in *order.OrderCreatedIn
				err := json.Unmarshal([]byte(d.Body), &in)
				if err != nil {
					log.Fatalf(err.Error())
				}

				oService := order.NewOrderService()
				oService.InsertOrder(in)
			}
		}
	}()

	fmt.Println("RabbitMQ Consumer started...")
	<-forever
}
