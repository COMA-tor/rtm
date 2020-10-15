package client

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type MqttClient struct {
	client mqtt.Client
}

type MqttConfiguration struct {
	Configuration
	Topic     string
	Qos       byte
	Retained  bool
	BrokerURI string
	ClientID  string
}

var NotMqttConfigurationError = errors.New("configuration is not of type MqttConfiguration")

func (m MqttClient) Publish(config Configuration, payload interface{}) error {
	switch config.(type) {
	case MqttConfiguration:
		mqttConfig := config.(MqttConfiguration)
		m.client.Publish(mqttConfig.Topic, mqttConfig.Qos, mqttConfig.Retained, payload)
		return nil
	default:
		return NotMqttConfigurationError
	}
}

func (m MqttClient) Subscribe(config Configuration, callback MessageHandler) error {
	switch config.(type) {
	case MqttConfiguration:
		mqttConfig := config.(MqttConfiguration)
		m.client.Subscribe(mqttConfig.Topic, mqttConfig.Qos, callback.(mqtt.MessageHandler))
		return nil
	default:
		return NotMqttConfigurationError
	}
}

func NewMqttClient(configuration MqttConfiguration) MqttClient {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(configuration.BrokerURI)
	opts.SetClientID(configuration.ClientID)
	return MqttClient{mqtt.NewClient(opts)}
}

func Testgo(client Client) {
	log.Println(client)
	log.Println(client.Publish(struct {}{}, "Ne doit pas passer"))
}
