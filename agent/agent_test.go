package agent

import (
	"errors"
	"time"

	"github.com/COMA-tor/rtm/sensor"
	"github.com/cucumber/godog"
)

var agent Agent
var sensorFeature sensor.Sensor
var measurements [][]byte

func iRunAnAgentThatUseIt() error {
	agent = NewMeasurementAgent(sensorFeature, func(m []byte) {
		measurements = append(measurements, m)
	})
	go agent.Run()
	return nil
}

func measurementsShouldBeCollected() error {
	select {
	case <-time.After(time.Second * 2):
		if len(measurements) > 0 {
			return nil
		}

		return errors.New("No measurements collected")
	}
}

func thatThereIsASensor() error {
	sensorFeature = sensor.NewSensor()
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I run an agent that use it$`, iRunAnAgentThatUseIt)
	s.Step(`^measurements should be collected$`, measurementsShouldBeCollected)
	s.Step(`^that there is a sensor$`, thatThereIsASensor)
}
