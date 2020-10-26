package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/COMA-tor/rtm/agent"
	"github.com/COMA-tor/rtm/sensor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffyaml"
)

type configuration struct {
	tickInterval time.Duration
	// sensorType   string
	topic      string
	qos        int
	brokerHost string
	brokerPort string
	clientID   string
}

const defaultTick = 10 * time.Second
const defaultClientID = ""
const defaultBrokerHost = "localhost"
const defaultBorkerPort = "1883"
const defaultTopic = "test"
const defaultQos = 0

var reload = make(chan int, 1)

func (config *configuration) init(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		tickInterval = flags.Duration("tick_interval", defaultTick, "Interval between two measurements")
		topic        = flags.String("topic", defaultTopic, "Topic where data should go")
		qos          = flags.Int("qos", defaultQos, "Quality Of Service for that agent")
		clientID     = flags.String("client_id", defaultClientID, "ID of the current agent")
		brokerHost   = flags.String("broker_host", defaultBrokerHost, "Host address of the broker")
		brokerPort   = flags.String("broker_port", defaultBorkerPort, "Port of the broker")
		_            = flags.String("config", "", "config file (optional)")
	)

	if err := ff.Parse(flags, args[1:],
		ff.WithEnvVarNoPrefix(),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ffyaml.Parser)); err != nil {
		return err
	}

	config.tickInterval = *tickInterval
	config.topic = *topic
	config.qos = *qos
	config.clientID = *clientID
	config.brokerHost = *brokerHost
	config.brokerPort = *brokerPort

	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	config := &configuration{}

	defer func() {
		signal.Stop(signalChannel)
		cancel()
	}()

	go func() {
		for {
			select {
			case signalReceived := <-signalChannel:
				switch signalReceived {
				case syscall.SIGINT, syscall.SIGTERM:
					signal.Stop(signalChannel)
					cancel()
					os.Exit(1)
				case syscall.SIGHUP:
					log.Println("Reload configuration")
					config.init(os.Args)
					reload <- 1
				}
			case <-ctx.Done():
				log.Println("Done")
				os.Exit(1)
			}
		}
	}()

	if err := run(ctx, config, os.Stdout); err != nil {
		log.SetOutput(os.Stderr)
		log.Fatalf("%s\n", err)
	}
}

func run(ctx context.Context, config *configuration, out io.Writer) error {
	config.init(os.Args)

	log.SetOutput(out)

	client, err := initClient(config)

	if err != nil {
		return err
	}

	defer func() {
		client.Disconnect(250)
	}()

	sensor := sensor.NewSensor()

	mqttAgent := agent.WithSensor(agent.EmptyAgent(), sensor, func(bytes []byte) {
		log.Println(string(bytes), config.topic, config.qos)
		client.Publish(config.topic, byte(config.qos), false, bytes)
	}, time.Second)

	go mqttAgent.Run(ctx)

	for {
		select {
		case <-reload:
			client, err = initClient(config)
			if err != nil {
				return err
			}
		}
	}
}

func initClient(config *configuration) (mqtt.Client, error) {
	options := mqtt.NewClientOptions().SetClientID(config.clientID).AddBroker(config.brokerHost + ":" + config.brokerPort)

	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}
