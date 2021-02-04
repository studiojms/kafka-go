package main

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "codepix",
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"T1"}
	c.SubscribeTopics(topics, nil)

	fmt.Println("kakfa is up and running")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println("->", string(msg.Value))
		}
	}
}
