package agent

import (
	"context"
	"time"

	"github.com/COMA-tor/rtm/sensor"
)

type Agent interface {
	Run(context.Context) error
}

type emptyAgent int

func (*emptyAgent) Run(context.Context) error {
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

func (s *sensorAgent) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(s.tickInterval):
			measurement := s.sensor.Value()
			s.handleMeasurement(measurement)
		}
	}

	return nil
}

func WithSensor(agent Agent, sensor sensor.Sensor, measurementHandler MeasurementHandler, tickInterval time.Duration) Agent {
	s := newSensorAgent(agent, sensor, measurementHandler, tickInterval)
	return &s
}

func newSensorAgent(agent Agent, sensor sensor.Sensor, measurementHandler MeasurementHandler, tickInterval time.Duration) sensorAgent {
	return sensorAgent{Agent: agent, sensor: sensor, handleMeasurement: measurementHandler, tickInterval: tickInterval}
}
