package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/buemura/btg-challenge/config"
	"github.com/buemura/btg-challenge/internal/modules/order"
	"github.com/buemura/btg-challenge/internal/queue"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {
	config.LoadEnv()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for {
		time.Sleep(time.Duration(rand.IntN(2)) * time.Second)

		body := &order.OrderCreatedIn{
			OrderID:    rand.IntN(1999) + 1,
			CustomerID: rand.IntN(9) + 1,
			Items:      generateItems(rand.IntN(3) + 1),
		}

		b, err := json.Marshal(body)
		if err != nil {
			failOnError(err, "Failed to marshall body")
		}

		err = ch.PublishWithContext(ctx,
			"",                        // exchange
			queue.ORDER_CREATED_QUEUE, // routing key
			false,                     // mandatory
			false,                     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        b,
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent message")
	}
}

func generateItems(qty int) []*order.Item {
	var items []*order.Item

	for i := 0; i < qty; i++ {
		item := &order.Item{
			Product:  fmt.Sprintf("Product %d", i),
			Quantity: rand.IntN(10) + 1,
			Price:    (rand.Float64() * 5) + 5,
		}
		items = append(items, item)
	}

	return items
}
