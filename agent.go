package main

import "time"

type Agent interface {
	Run() error
}

type emptyAgent int

func (*emptyAgent) Run() error {
	return nil
}

func NewAgent() Agent {
	return new(emptyAgent)
}

type DefaultAgent struct {
	businessFunction func()
}

func (agent *DefaultAgent) Run() error {
	for {
		select {
		case <-time.Tick(time.Second):
			agent.businessFunction()
		}
	}
}

func WithHandler(agent Agent, handler func()) Agent {
	return &DefaultAgent{businessFunction: handler}
}

type MeasurementAgent struct {
	DefaultAgent
}

func NewMeasurementAgent(sensor Sensor, measurementHandler func([]byte)) Agent {
	agent := WithHandler(NewAgent(), func() {
		value := sensor.Value()
		measurementHandler(value)
	})
	return agent
}
