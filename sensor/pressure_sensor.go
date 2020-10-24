package sensor

import (
	"fmt"
	"time"

	"github.com/COMA-tor/rtm/sensor/generator"
)

// pressureValue returns a radom value generated using Perlin generator.
func pressureValue() []byte {
	value := generator.TemperatureGenerator(time.Now().Unix())

	return []byte(fmt.Sprint(value))
}

// PressureSensor is a sensor that provides pressure value.
func PressureSensor() Sensor {
	return WithCustomValue(EmptySensor(), pressureValue)
}
