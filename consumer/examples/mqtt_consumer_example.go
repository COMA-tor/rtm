package main

import (
	"log"

	"github.com/COMA-tor/rtm/consumer"
)

func main() {
	mqttConsumer := consumer.NewMqttConsumer(
		"airport/#",
		"localhost",
		"1883",
		func(bytes []byte) {
			log.Printf("Data received: %v", bytes)
		},
	)
	mqttConsumer.Run()
}
