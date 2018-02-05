package util

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

//RegisterConsumer Registers an amqp consumer on the given queue
func RegisterConsumer(queue string, consume func([]byte) error) {
	conn, ch := declareExchangeAndQueue(queue)
	defer ch.Close()
	defer conn.Close()
	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	exitOnErr("Failed to register a consumer", err)

	forever := make(chan bool)
	go checkRabbitCon(ch, queue)
	go func() {
		for d := range msgs {
			err := consume(d.Body)
			if err == nil {
				d.Ack(false)
			} else {
				d.Reject(true)
				fmt.Println(errors.Wrap(err, ""))
			}
		}
	}()

	fmt.Println("Waiting for messages on " + queue)
	<-forever
}

//QueueInspect counts number of messages in the queue
func QueueInspect(name string) (int, error) {
	conn, ch := declareExchangeAndQueue(name)
	queue, err := ch.QueueInspect(name)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error inspecting queue"))
		return 0, err
	}
	ch.Close()
	conn.Close()
	return queue.Messages, nil
}

func checkRabbitCon(ch *amqp.Channel, qName string) {

	ticker := time.NewTicker(time.Second * 30)
	for range ticker.C {
		_, err := ch.QueueInspect(qName)
		if err != nil {
			fmt.Println(errors.Wrap(err, "Lost Rabbit Connection: "))
			os.Exit(1)
		}
	}
}

func declareExchangeAndQueue(queue string) (*amqp.Connection, *amqp.Channel) {
	amqpAddress := GetConfigValue("amqpAddress")
	conn, err := amqp.Dial(amqpAddress)
	exitOnErr("Failed to connect to RabbitMQ "+amqpAddress, err)

	ch, err := conn.Channel()
	exitOnErr("Failed to open a channel", err)

	err = ch.Qos(1, 0, true)
	exitOnErr("Failed to open a channel", err)
	err = ch.ExchangeDeclare(queue, "fanout", true, false, false, false, nil)
	exitOnErr("Failed to declare exchange", err)

	q, err := ch.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	exitOnErr("Failed to declare a queue", err)
	err = ch.QueueBind(q.Name, "", queue, false, nil)
	exitOnErr("Failed to bind exchange to queue", err)

	return conn, ch
}

//Publish publishes a message on given queue
func Publish(msg interface{}, queue string) error {
	conn, ch := declareExchangeAndQueue(queue)
	defer ch.Close()
	defer conn.Close()
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "Couldn't unmarshal")
	}
	err = ch.Publish(
		queue, // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonBytes,
		})
	//log.Printf(" [x] Sent %s", *body)
	if err != nil {
		return errors.Wrap(err, "Error publishing")
	}
	return nil
}

func exitOnErr(msg string, err error) {
	if err != nil {
		fmt.Println(errors.Wrap(err, msg))
		os.Exit(-1)
	}
}
