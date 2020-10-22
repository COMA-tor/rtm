package sensor

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
