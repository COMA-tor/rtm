package src

import "github.com/COMA-tor/rtm/mqttclient"

type Consumer interface {
	// Executes the consumer
	Run()
}

// mqttConsumer implements Consumer
type MqttConsumer struct {
	BrokerIP string
	BrokerPort int
	ClientID string
}

func (m *MqttConsumer) Run() {
	broker := mqttclient.MqttBroker{IP: m.BrokerIP, Port: m.BrokerPort}

	client := broker.Connect(m.ClientID)
	defer client.Disconnect(250)
}


