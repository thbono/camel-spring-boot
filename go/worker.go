package main

import (
	"log"
	"os"
	"database/sql"

	"github.com/streadway/amqp"
	_ "gopkg.in/rana/ora.v4"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("AMQP_QUEUE_NAME"), // name
		true, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1, // prefetch count
		0, // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			db, err := sql.Open("ora", "consignetflexprod/consignetflexprod@10.13.170.251:1521/orcl")
			failOnError(err, "Failed to connect to dataabse")
			defer db.Close()

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}