package main

import (
	"github.com/COMA-tor/rtm/consumer/src"
)

const BrokerIP = "localhost"
const BrokerPort = 1883
const ConsID = "CONSUMER-001"

func main() {
	cons := src.MqttConsumer{BrokerIP: BrokerIP, BrokerPort: BrokerPort, ClientID: ConsID}
	cons.Run()
}
