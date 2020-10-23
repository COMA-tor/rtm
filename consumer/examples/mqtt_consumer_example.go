package main

import (
	"github.com/COMA-tor/rtm/consumer"
	"log"
)

func main() {
	mqttConsumer := consumer.NewMqttConsumer(
		func(bytes []byte) {
			log.Printf("Data received: %v", bytes)
		},
	)
	mqttConsumer.Run()
}
