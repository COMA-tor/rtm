package src

import "testing"

const BrokerIP = "localhost"
const BrokerPort = 1883
const ConsID = "CONSUMER-001"

func TestMqttConsumer_Run(t *testing.T) {
	t.Run("Connect with same ID", func(t *testing.T) {
		consumer1 := MqttConsumer{BrokerIP, BrokerPort, ConsID}
		consumer1.Run()

		// consumer1 connection is closed
		consumer2 := MqttConsumer{BrokerIP, BrokerPort, ConsID}
		consumer2.Run()
	})
}