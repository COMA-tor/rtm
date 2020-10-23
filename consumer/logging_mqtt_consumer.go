package consumer

import (
	"log"
	"os"
)

const brokerHost = "localhost"
const brokerPort = "1883"
const clientId = "CLIENT-001"

type MqttToLogConsumer struct {
	MqttConsumer
}

func newLogHandler(file string) func([]byte) {
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

func NewMqttToLogConsumer(logFile string) MqttToLogConsumer {
	return MqttToLogConsumer{
		MqttConsumer: NewMqttConsumer("test", newLogHandler(logFile)),
	}
}
