package sensor

import (
	"fmt"

	"github.com/cucumber/godog"
)

var sensor Sensor
var value []byte

func iReadTheSensorValue() error {
	value = sensor.Value()
	return nil
}

func theValueShouldBeNil() error {
	if value != nil {
		return fmt.Errorf("expected nil value, but got %v", value)
	}

	return nil
}

func thereIsAnEmptySensor() error {
	sensor = EmptySensor()
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^i read the sensor value$`, iReadTheSensorValue)
	s.Step(`^the value should be nil$`, theValueShouldBeNil)
	s.Step(`^there is an empty sensor$`, thereIsAnEmptySensor)
}
