package sensor

import (
	"log"
	"text/template"
	"time"
)

/*
opt := NewClientOptions()
opt.AddBroker("tcp://172.16.128.80:1883")

client := NewClient(opt)
defer client.Disconnect(250)

if token := client.Connect(); token.Wait() && token.Error() != nil {
	log.Fatal(token.Error())
	return
}
*/

type config struct {
	interval        time.Duration
	min, max        float64
	logFormat, unit string
}

func Run(cfg config, generator GeneratorFunc) {
	for {
		select {
		case <-time.Tick(cfg.interval * time.Second):
			unix := time.Now().Unix()
			value := generator(unix)
			if cfg.logFormat == "" {
				cfg.logFormat = "%.4f%s"
			}
			tmpl, err := template.New("log").Parse(cfg.logFormat)
			// client.Publish(topic, 0, false, fmt.Sprintf("%d|%.4f", unix, data))
			data := struct {
				Value float64
				Unit  string
			}{value, cfg.unit}
			err = tmpl.Execute(log.Writer(), data)
			if err != nil {
				log.Fatal("Error executing log template at sensor.run:", err)
			}
		}
	}
}
