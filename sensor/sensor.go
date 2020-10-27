package sensor

import "log"

// A Sensor carries technical informations and exposes a value that can be read
// in order to be used for computations.
type Sensor interface {
	// Value returns the current value of the sensor.
	Value() []byte
}

// An emptySensor always has nil values. It is not struct{}, since vars of this
// type must have distinct addresses.
type emptySensor int

func (*emptySensor) Value() []byte {
	return nil
}

func NewSensor() Sensor {
	return new(emptySensor)
}

// EmptySensor returns a non-nil, empty Sensor. It always has nil values. It is
// typically used by initialization and tests.
func EmptySensor() Sensor {
	return new(emptySensor)
}

// A ValueMethod tells the sensor how the value is read.
type ValueMethod func() []byte

// WithCustomValue returns a copy of sensor that use the callback parameter to
// read the sensor's value.
func WithCustomValue(sensor Sensor, callback ValueMethod) Sensor {
	if sensor == nil {
		log.Panic("cannot create sensor from nil one")
	}

	customizedSensor := newCustomSensor(sensor, callback)

	return &customizedSensor
}

// newCustomSensor returns an initialized customSensor.
func newCustomSensor(sensor Sensor, callback ValueMethod) customSensor {
	return customSensor{Sensor: sensor, valueMethod: callback}
}

// A customSensor is a sensor with a defined value method.
type customSensor struct {
	Sensor

	valueMethod ValueMethod
}

func (c *customSensor) Value() []byte {
	return c.valueMethod()
}
