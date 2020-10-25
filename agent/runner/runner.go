package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/COMA-tor/rtm/agent"
	"github.com/COMA-tor/rtm/sensor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/namsral/flag"
)

type configuration struct {
	tickInterval time.Duration
	// sensorType   string
	// airportIATA  string
	brokerHost string
	brokerPort string
	clientID   string
}

const defaultTick = 10 * time.Second
const defaultClientID = ""
const defaultBrokerHost = "localhost"
const defaultBorkerPort = "1883"

func (config *configuration) init(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.String(flag.DefaultConfigFlagname, "", "Path to config file")

	var (
		tickInterval = flags.Duration("tick_interval", defaultTick, "Interval between two measurements")
		clientID     = flags.String("client_id", defaultClientID, "ID of the current agent")
		brokerHost   = flags.String("broker_host", defaultBrokerHost, "Host address of the broker")
		brokerPort   = flags.String("broker_port", defaultBorkerPort, "Port of the broker")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	config.tickInterval = *tickInterval
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

	options := mqtt.NewClientOptions().SetClientID(config.clientID).AddBroker(config.brokerHost + ":" + config.brokerPort)

	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	defer func() {
		client.Disconnect(250)
	}()

	sensor := sensor.NewSensor()

	mqtt_agent := agent.WithSensor(agent.EmptyAgent(), sensor, func(bytes []byte) {
		log.Println(string(bytes))
		client.Publish("test", 1, false, bytes)
	}, time.Second)

	mqtt_agent.Run()

	return nil
}
