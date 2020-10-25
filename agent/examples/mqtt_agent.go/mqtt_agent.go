package main

import (
	"log"
	"time"

	"github.com/COMA-tor/rtm/agent"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const brokerHost = "localhost"
const brokerPort = "1883"
const clientId = "CLIENT-002"

type MockSensor struct{}

func (MockSensor) Value() []byte {
	return []byte(time.Now().Format(time.Stamp))
}

func main() {
	log.Printf("Trying to connect (%s, %s)", brokerHost, brokerPort)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerHost + ":" + brokerPort)
	opts.SetClientID(clientId)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); !token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println("client1 is connected :", client.IsConnected())

	mqttAgent := agent.WithSensor(agent.EmptyAgent(), MockSensor{}, func(bytes []byte) {
		log.Println(string(bytes))
		client.Publish("test", 1, false, bytes)
	}, time.Second)

	mqttAgent.Run()
}
