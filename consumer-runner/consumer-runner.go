package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/COMA-tor/rtm/consumer"
	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffyaml"
)

type configuration struct {
	// qos        int
	brokerHost string
	brokerPort string
	redisHost  string
	redisPort  string
	clientID   string
}

const defaultClientID = ""
const defaultBrokerHost = "localhost"
const defaultBrokerPort = "1883"
const defaultRedisHost = "localhost"
const defaultRedisPort = "6379"

// const defaultQos = 0

func (config *configuration) init(args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		// qos          = flags.Int("qos", defaultQos, "Quality Of Service for that agent")
		clientID   = flags.String("client_id", defaultClientID, "ID of the current agent")
		brokerHost = flags.String("broker_host", defaultBrokerHost, "Host address of the broker")
		brokerPort = flags.String("broker_port", defaultBrokerPort, "Port of the broker")
		redisHost  = flags.String("redis_host", defaultBrokerHost, "Host address of the broker")
		redisPort  = flags.String("redis_port", defaultRedisPort, "Port of the broker")
		_          = flags.String("config", "", "config file (optional)")
	)

	if err := ff.Parse(flags, args[1:],
		ff.WithEnvVarNoPrefix(),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ffyaml.Parser)); err != nil {
		return err
	}

	config.clientID = *clientID
	config.brokerHost = *brokerHost
	config.brokerPort = *brokerPort
	config.redisHost = *redisHost
	config.redisPort = *redisPort

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
				case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP:
					signal.Stop(signalChannel)
					cancel()
					os.Exit(1)
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

	mqttToRedisConsumer := consumer.NewMqttToRedisConsumer(
		config.redisHost,
		config.redisPort,
		config.brokerHost,
		config.brokerPort,
		config.clientID,
	)

	return mqttToRedisConsumer.Run()
}
