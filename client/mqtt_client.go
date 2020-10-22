package client

import (
	"errors"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttClient struct {
	client        mqtt.Client
	configuration MqttConfiguration
}

// MqttConfiguration define the configuration used by a mqttClient
type MqttConfiguration struct {
	Configuration
	Topic     string
	Qos       byte
	Retained  bool
	BrokerURI string
	ClientID  string
}

var errorNotMqttConfiguration = errors.New("configuration is not of type MqttConfiguration")

func (m mqttClient) Publish(config Configuration, payload Payload) error {
	switch config.(type) {
	case MqttConfiguration:
		mqttConfig := config.(MqttConfiguration)
		m.client.Publish(mqttConfig.Topic, mqttConfig.Qos, mqttConfig.Retained, payload)
		return nil
	default:
		return errorNotMqttConfiguration
	}
}

func (m mqttClient) Subscribe(config Configuration, callback MessageHandler) error {
	switch config.(type) {
	case MqttConfiguration:
		mqttConfig := config.(MqttConfiguration)
		m.client.Subscribe(mqttConfig.Topic, mqttConfig.Qos, callback.(mqtt.MessageHandler))
		return nil
	default:
		return errorNotMqttConfiguration
	}
}

// NewMqttClient will create an MQTT client with all the specified options
func NewMqttClient(configuration *MqttConfiguration) Client {
	client := &mqttClient{}

	if configuration != nil {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(configuration.BrokerURI)
		opts.SetClientID(configuration.ClientID)

		client.client = mqtt.NewClient(opts)
	}

	return client
}
