package agent

import (
	"time"

	"github.com/COMA-tor/rtm/sensor"
)

type Agent interface {
	Run() error
}

type emptyAgent int

func (*emptyAgent) Run() error {
	return nil
}

func EmptyAgent() Agent {
	return new(emptyAgent)
}

type MeasurementHandler func([]byte)

type sensorAgent struct {
	Agent
	sensor            sensor.Sensor
	handleMeasurement MeasurementHandler
	tickInterval      time.Duration
}

func (s sensorAgent) Run() error {
	for {
		select {
		case <-time.Tick(s.tickInterval):
			measurement := s.sensor.Value()
			s.handleMeasurement(measurement)
		}
	}

	return nil
}

func WithSensor(agent Agent, sensor sensor.Sensor, measurementHandler MeasurementHandler, tickInterval time.Duration) Agent {
	return newSensorAgent(agent, sensor, measurementHandler, tickInterval)
}

func newSensorAgent(agent Agent, sensor sensor.Sensor, measurementHandler MeasurementHandler, tickInterval time.Duration) sensorAgent {
	return sensorAgent{Agent: agent, sensor: sensor, handleMeasurement: measurementHandler, tickInterval: tickInterval}
}
