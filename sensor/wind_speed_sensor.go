package sensor

import (
	"fmt"
	"time"

	"github.com/COMA-tor/rtm/sensor/generator"
)

// windSpeedValue returns a radom value generated using Perlin generator.
func windSpeedValue() []byte {
	value := generator.WindSpeedGenerator(time.Now().Unix())

	return []byte(fmt.Sprint(value))
}

// WindSpeedSensor is a sensor that provides wind speed value.
func WindSpeedSensor() Sensor {
	return WithCustomValue(EmptySensor(), windSpeedValue)
}
