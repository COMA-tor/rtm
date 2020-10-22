package consumer

import (
	"errors"
	"github.com/cucumber/godog"
	"time"
)

var consumer Consumer
var measurements [][]byte

func gen() <-chan []byte {
	out := make(chan []byte)
	go func() {
		for _, n := range []string{"hello"} {
			out <- []byte(n)
		}
		close(out)
	}()
	return out
}

func handleData(data []byte) {
	measurements = append(measurements, data)
}

func theConsumerIsRunning() error {
	go consumer.Run()
	return nil
}

func itShouldHandleData() error {
	select {
	case <-time.After(time.Millisecond * 10):
		if len(measurements) > 0 {
			return nil
		}
		return errors.New("no measurements received")
	}
}

func thereIsAConsumer() error {
	consumer = WithDataHandler(NewConsumer(), gen, handleData)
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^It should handle data$`, itShouldHandleData)
	s.Step(`^The consumer is running$`, theConsumerIsRunning)
	s.Step(`^There is a consumer$`, thereIsAConsumer)
}
