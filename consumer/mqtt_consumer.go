package consumer

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type MqttConsumer struct {
	DefaultConsumer
	client mqtt.Client
}

func NewMqttConsumer(topic string, handleData func(bytes []byte)) MqttConsumer {
	log.Printf("Trying to connect (%s, %s)", brokerHost, brokerPort)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerHost + ":" + brokerPort)
	opts.SetClientID(clientId)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); !token.WaitTimeout(3 * time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Println("client is connected :",client.IsConnected())

	return MqttConsumer{
		DefaultConsumer: DefaultConsumer{
			listenData: func() <-chan []byte {
				out := make(chan []byte)
				client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
					out <- topicAndDataToBytes(message.Topic(), message.Payload())
				})
				return out
			},
			handleData: handleData,
		},
		client: client,
	}
}

func topicAndDataToBytes(topic string, payload []byte) []byte {
	return []byte(
		fmt.Sprintf(
			"%s</payload>%v",
			topic,
			payload,
		),
	)
}
