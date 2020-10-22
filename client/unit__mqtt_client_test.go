package client

import "testing"

func init() {}

func TestSimpleMqttClient(t *testing.T) {
	configuration := &MqttConfiguration{}
	client := NewMqttClient(configuration).(*mqttClient)

	if client.client == nil {
		t.Fatal("client is nil")
	}
}
