package sensor

import (
	"fmt"
	"time"

	"github.com/COMA-tor/rtm/sensor/generator"
)

// temperatureValue returns a radom value generated using Perlin generator.
func temperatureValue() []byte {
	value := generator.TemperatureGenerator(time.Now().Unix())

	return []byte(fmt.Sprint(value))
}

// A TemperatureSensor is a sensor that provides temperature value.
func TemperatureSensor() Sensor {
	return WithCustomValue(EmptySensor(), temperatureValue)
}
