package main

import "github.com/COMA-tor/rtm/consumer"

func main() {
	mqttConsumer := consumer.NewMqttConsumer()
	mqttConsumer.Run()
}
