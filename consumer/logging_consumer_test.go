package consumer

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"io/ioutil"
	"os"
	"time"
)

var loggingConsumer MqttToLogConsumer

const logFile = "examples/data.log"

func itShouldWriteReceivedDataInFile() error {
	select {
	case <-time.After(time.Millisecond * 10):
		data, err := ioutil.ReadFile(logFile)
		if err != nil {
			return errors.Unwrap(err)
		}
		if len(data) > 0 {
			return nil
		}
		return errors.New("file is empty")
	}
}

func theLogConsumerIsRunning() error {
	go loggingConsumer.Run()
	return nil
}

func thereIsALogConsumer() error {
	loggingConsumer = NewMqttToLogConsumer(logFile)
	loggingConsumer.listenData = gen
	return nil
}

func lineShouldBeWrittenInFile(nbLines int) error {
	select {
	case <-time.After(time.Millisecond * 10):
		file, err := os.Open(logFile)
		if err != nil {
			return errors.Unwrap(err)
		}

		scanner := bufio.NewScanner(file)
		countLines := 0
		for scanner.Scan() {
			countLines++
		}

		if err := scanner.Err(); err != nil {
			return errors.Unwrap(err)
		}

		if nbLines != countLines {
			return errors.New(fmt.Sprintf("file was expected to have %d lines, actual is %d", nbLines, countLines))
		}
		return nil
	}
}

func theLogConsumerReceiveSliceOfBytes(nbLines int) error {
	loggingConsumer = NewMqttToLogConsumer(logFile)
	loggingConsumer.listenData = func() <-chan []byte {
		out := make(chan []byte)
		go func() {
			for _, n := range make([]int, nbLines) {
				out <- []byte(fmt.Sprintf("Line n°%d", n))
			}
			close(out)
		}()
		return out
	}
	return nil
}


func LogConsumerFeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(pickle *messages.Pickle) {
		_ = os.Remove(logFile)
	})
	s.Step(`^It should write received data in file$`, itShouldWriteReceivedDataInFile)
	s.Step(`^The log consumer is running$`, theLogConsumerIsRunning)
	s.Step(`^There is a log consumer$`, thereIsALogConsumer)
	s.Step(`^(\d+) line should be written in file$`, lineShouldBeWrittenInFile)
	s.Step(`^The log consumer receive (\d+) slice of bytes$`, theLogConsumerReceiveSliceOfBytes)
}
