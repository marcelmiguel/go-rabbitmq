package rabbitmq

import (
	"fmt"
	"testing"
	"time"
)

const URITEST = "amqp://neo:rocamora56@amqpcloudtest.neogrup.com:5672/"
const QUEUE = "NGRPMMC4.shops.neogrup.com"

func TestConsumers(t *testing.T) {

	consumers, err := Consumers(URITEST, QUEUE)
	if err != nil {
		t.Errorf("Cannot publish %v\n", err)
	}
	if consumers != 0 {
		t.Errorf("There is no conenction ! %d\n", consumers)
	}

	config := Config{}
	consumer, err := NewConsumer(URITEST, config)
	if err != nil {
		t.Errorf("Failed initialize %v\n", err)
	}

	err = consumer.StartConsuming(
		func(d Delivery) Action {
			fmt.Printf("%s Recieved: %s\n", QUEUE, time.Now().Format("150405"))
			return Ack
		},
		QUEUE,
		[]string{},
		WithConsumeOptionsConsumerName("consumerName"),
	)
	if err != nil {
		t.Errorf("Cannot start consuming %v\n", err)
	}

	consumers, err = Consumers(URITEST, QUEUE)
	if err != nil {
		t.Errorf("Cannot publish %v\n", err)
	}
	if consumers != 1 {
		t.Errorf("There is no conenction ! %d\n", consumers)
	}

	consumer.Close()

}
