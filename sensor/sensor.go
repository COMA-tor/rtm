package sensor

import (
	"log"
	"text/template"
	"time"
)

// Configuration used by the sensor.
type Configuration struct {
	Interval        time.Duration
	Min, Max        float64
	LogFormat, Unit string
}

// Run the sensor with the given configuration and that use the given generator
// function.
func Run(cfg Configuration, generator GeneratorFunc) {
	for {
		select {
		case <-time.Tick(cfg.Interval * time.Second):
			unix := time.Now().Unix()
			value := generator(unix)
			if cfg.LogFormat == "" {
				cfg.LogFormat = "%.4f%s"
			}
			tmpl, err := template.New("log").Parse(cfg.LogFormat)

			if err != nil {
				log.Fatal("Error creating log template at sensor.run:", err)
			}

			data := struct {
				Value float64
				Unit  string
			}{value, cfg.Unit}

			err = tmpl.Execute(log.Writer(), data)

			if err != nil {
				log.Fatal("Error executing log template at sensor.run:", err)
			}
		}
	}
}
