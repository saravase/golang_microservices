package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "pkc-41p56.asia-south1.gcp.confluent.cloud:9092",
		"security.protocol": "SASL_SSL",
		"sasl.username":     "US5QOCFZJU5O3DRK",
		"sasl.password":     "fytBrrleTzNOjd1kUUG10jlXeWUiuaBEsd1Lt9wq18TPdLcz2xhvsx0OCQ4Eo6q5",
		"sasl.mechanism":    "PLAIN",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{"default"}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	consumer.Close()

	//ambassadorMessage := []byte(fmt.Sprintf("You earned $%f from the link #%s", ambassadorRevenue, order.Code))

	//	smtp.SendMail("host.docker.internal:1025", nil, "no-reply@email.com", []string{order.AmbassadorEmail}, ambassadorMessage)

	//	adminMessage := []byte(fmt.Sprintf("Order #%d with a total of $%f has been completed", order.Id, adminRevenue))

	//	smtp.SendMail("host.docker.internal:1025", nil, "no-reply@email.com", []string{"admin@admin.com"}, adminMessage)

}
