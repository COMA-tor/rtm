package consumer

import (
	"log"
	"os"
)

const brokerHost = "localhost"
const brokerPort = "1883"
const clientId = "CLIENT-001"

type LocalMqttConsumer struct {
	MqttConsumer
}

func newLocalHandler(file string) func([]byte) {
	return func(data []byte) {
		if len(data) > 0 {
			f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
				f.Close()
				return
			}
			_, err = f.WriteString(string(data) + "\n")
			if err != nil {
				log.Fatal(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
				f.Close()
				return
			}
		}
	}
}

func NewLocalMqttConsumer(localFile string) LocalMqttConsumer {
	return LocalMqttConsumer{
		MqttConsumer: NewMqttConsumer(newLocalHandler(localFile)),
	}
}
