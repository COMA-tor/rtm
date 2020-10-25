package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/COMA-tor/rtm/sensor"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

var (
	ctx                 context.Context
	agent               Agent
	sensorFeature       sensor.Sensor
	measurements        [][]byte
	measurementInterval time.Duration
	measurementsCount   int
)

func thereIsASensor() error {
	sensorFeature = sensor.NewSensor()
	return nil
}

func thereIsAMeasurementIntervalOfMilliseconds(expectedInterval int) error {
	measurementInterval = time.Duration(expectedInterval) * time.Millisecond
	return nil
}

func thereIsAnAgentThatUseIt() error {
	agent = WithSensor(EmptyAgent(), sensorFeature, func(b []byte) {
		measurements = append(measurements, b)
	}, measurementInterval)

	return nil
}

func iRunTheAgentForMilliseconds(expectedRunDuration int) error {
	ctx, _ := context.WithTimeout(ctx, time.Duration(expectedRunDuration)*time.Millisecond)

	agent.Run(ctx)

	return nil
}

func thereShouldBeMeasurementsCollected(expectedMeasurementsCount int) error {
	if len(measurements) != expectedMeasurementsCount {
		return fmt.Errorf("expected %v measurements, got %v", expectedMeasurementsCount, len(measurements))
	}

	measurementsCount = expectedMeasurementsCount

	return nil
}

func noMoreShouldBeCollected() error {
	select {
	case <-time.After(time.Second):
		if measurementsCount != len(measurements) {
			return fmt.Errorf("expected %v measurements, got %v", measurementsCount, len(measurements))
		}
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(m *messages.Pickle) {
		ctx = context.Background()
		agent = EmptyAgent()
		sensorFeature = sensor.NewSensor()
		measurements = make([][]byte, 0)
		measurementInterval = time.Second

	})

	s.Step(`^there is a sensor$`, thereIsASensor)
	s.Step(`^there is a measurement interval of (\d+) milliseconds$`, thereIsAMeasurementIntervalOfMilliseconds)
	s.Step(`^there is an agent that use it$`, thereIsAnAgentThatUseIt)

	s.Step(`^I run the agent for (\d+) milliseconds$`, iRunTheAgentForMilliseconds)

	s.Step(`^there should be (\d+) measurements collected$`, thereShouldBeMeasurementsCollected)
	s.Step(`^no more should be collected$`, noMoreShouldBeCollected)
}
