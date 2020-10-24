package sensor

import (
	"fmt"
	"time"

	"github.com/COMA-tor/rtm/sensor/generator"
)

func temperatureValue() []byte {
	value := generator.TemperatureGenerator(time.Now().Unix())

	return []byte(fmt.Sprint(value))
}

func TemperatureSensor() Sensor {
	return WithCustomValue(EmptySensor(), temperatureValue)
}
