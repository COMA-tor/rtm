package mqttclient

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type MqttBroker struct {
	IP     string
	Port   int
}

type mqttOptions struct {
	BrokerIP 	string
	BrokerPort 	int
	ClientID 	string
}

func (m *MqttBroker) Connect(clientID string) mqtt.Client {
	fmt.Printf("Try to connect to broker '%s' with port '%d' ...\n", m.IP, m.Port)

	opts := mqttOptions{
		BrokerIP: m.IP,
		BrokerPort: m.Port,
		ClientID: clientID,
	}

	client := mqtt.NewClient(opts.ClientOptions())

	if token := client.Connect(); !token.WaitTimeout(3 * time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}

	fmt.Printf("Connected with ID '%s'\n", clientID)
	return client
}

func (o *mqttOptions) ClientOptions() *mqtt.ClientOptions {
	res := mqtt.NewClientOptions()

	res.AddBroker(
		fmt.Sprintf(
			"tcp://%s:%d",
			o.BrokerIP,
			o.BrokerPort,
		),
	)
	res.SetClientID(o.ClientID)

	return res
}
