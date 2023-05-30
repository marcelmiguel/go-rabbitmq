package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumers(amqpURL string, queueName string) (int, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return 0, err
	}
	defer ch.Close()

	queueInfo, err := ch.QueueInspect(queueName)
	if err != nil {
		return 0, err
	}

	return queueInfo.Consumers, nil
}
