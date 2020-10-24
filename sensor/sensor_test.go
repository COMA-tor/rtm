package sensor

import (
	"fmt"
	"strconv"

	"github.com/cucumber/godog"
)

var sensor Sensor
var value []byte
var callback func() []byte

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

func theCallbackFunctionIsDefinedForTheSensor() error {
	sensor = WithCustomValue(sensor, callback)
	return nil
}

func theValueShouldBe(expectedInt int) error {
	expectedValue := strconv.Itoa(expectedInt)

	if string(value) != expectedValue {
		return fmt.Errorf("expected %v value, but got %v", expectedValue, string(value))
	}

	return nil
}

func thereIsACallbackFunctionThatReturn(returnValue int) error {
	callback = func() []byte {
		return []byte(strconv.Itoa(returnValue))
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^i read the sensor value$`, iReadTheSensorValue)
	s.Step(`^the value should be nil$`, theValueShouldBeNil)
	s.Step(`^there is an empty sensor$`, thereIsAnEmptySensor)

	s.Step(`^the callback function is defined for the sensor$`, theCallbackFunctionIsDefinedForTheSensor)
	s.Step(`^the value should be (\d+)$`, theValueShouldBe)
	s.Step(`^there is a callback function that return (\d+)$`, thereIsACallbackFunctionThatReturn)
}
