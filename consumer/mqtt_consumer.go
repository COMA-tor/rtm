package consumer

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type MqttConsumer struct {
	DefaultConsumer
	client mqtt.Client
}

const brokerHost = "localhost"
const brokerPort = "1883"
const clientId = "CLIENT-001"

func NewMqttConsumer() MqttConsumer {
	log.Printf("Trying to connect (%s, %s)", brokerHost, brokerPort)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerHost + ":" + brokerPort)
	opts.SetClientID(clientId)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); !token.WaitTimeout(3 * time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println("client1 is connected :",client.IsConnected())

	return MqttConsumer{
		DefaultConsumer: DefaultConsumer{
			listenData: func() <-chan []byte {
				out := make(chan []byte)
				client.Subscribe("test", 0, func(client mqtt.Client, message mqtt.Message) {
					out <- message.Payload()
				})
				return out
			},
			handleData: func(bytes []byte) {
				log.Printf("Data received: %v", bytes)
			},
		},
		client: client,
	}
}
