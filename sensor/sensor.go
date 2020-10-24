package sensor

import "log"

type Sensor interface {
	Value() []byte
}

type emptySensor int

func (*emptySensor) Value() []byte {
	return nil
}

func NewSensor() Sensor {
	return new(emptySensor)
}

func EmptySensor() Sensor {
	return new(emptySensor)
}

func WithCustomValue(sensor Sensor, callback ValueMethod) Sensor {
	if sensor == nil {
		log.Panic("cannot create sensor from nil one")
	}

	customizedSensor := newCustomSensor(sensor, callback)

	return &customizedSensor
}

func newCustomSensor(sensor Sensor, callback ValueMethod) customSensor {
	return customSensor{Sensor: sensor, valueMethod: callback}
}

type ValueMethod func() []byte

type customSensor struct {
	Sensor

	valueMethod ValueMethod
}

func (c *customSensor) Value() []byte {
	return c.valueMethod()
}
