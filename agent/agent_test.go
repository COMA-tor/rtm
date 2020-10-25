package agent

import (
	"fmt"
	"time"

	"github.com/COMA-tor/rtm/sensor"
	"github.com/cucumber/godog"
)

var (
	agent               Agent
	sensorFeature       sensor.Sensor
	measurements        [][]byte
	measurementInterval time.Duration
)

func iRunAnAgentThatUseIt() error {
	return godog.ErrPending
}

func iRunTheAgentForMilliseconds(expectedRunDuration int) error {
	go agent.Run()

	select {
	case <-time.After(time.Duration(expectedRunDuration) * time.Millisecond):
		return nil
	}
}

func thereIsASensor() error {
	sensorFeature = sensor.NewSensor()
	return nil
}

func thereIsAnAgentThatUseIt() error {
	agent = WithSensor(EmptyAgent(), sensorFeature, func(b []byte) {
		measurements = append(measurements, b)
	}, measurementInterval)

	return nil
}

func thereShouldBeMeasurementsCollected(expectedMeasurementsCount int) error {
	if len(measurements) != expectedMeasurementsCount {
		return fmt.Errorf("expected %v measurements, got %v", expectedMeasurementsCount, len(measurements))
	}

	return nil
}

func thereIsAMeasurementIntervalOfMilliseconds(expectedInterval int) error {
	measurementInterval = time.Duration(expectedInterval) * time.Millisecond
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I run an agent that use it$`, iRunAnAgentThatUseIt)
	s.Step(`^I run the agent for (\d+) milliseconds$`, iRunTheAgentForMilliseconds)
	s.Step(`^there is a sensor$`, thereIsASensor)
	s.Step(`^there is an agent that use it$`, thereIsAnAgentThatUseIt)
	s.Step(`^there should be (\d+) measurements collected$`, thereShouldBeMeasurementsCollected)
	s.Step(`^there is a measurement interval of (\d+) milliseconds$`, thereIsAMeasurementIntervalOfMilliseconds)
}
