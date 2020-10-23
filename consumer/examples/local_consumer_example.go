package main

import "github.com/COMA-tor/rtm/consumer"

var localFile = "consumer/examples/mqtt_data_example.log"

func myListenData() <-chan []byte {
	out := make(chan []byte)
	go func() {
		for _, n := range []string{"hello", "world", "!"} {
			out <- []byte(n)
		}
		close(out)
	}()
	return out
}

func main() {
	localMqttConsumer := consumer.NewLocalMqttConsumer(localFile)
	localMqttConsumer.Run()
}
