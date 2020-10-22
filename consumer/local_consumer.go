package consumer

import (
	"log"
	"os"
)

type LocalConsumer struct {
	DefaultConsumer
}

func newLocalHandler(file string) func([]byte) {
	return func(data []byte) {
		if len(data) > 0 {
			f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
				f.Close()
				return
			}
			_, err = f.WriteString(string(data) + "\n")
			if err != nil {
				log.Fatal(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				log.Fatal(err)
				f.Close()
				return
			}
		}
	}
}

func NewLocalConsumer(localFile string, listenData func() <-chan []byte) LocalConsumer {
	defaultConsumer := WithDataHandler(NewConsumer(), listenData, newLocalHandler(localFile))
	return LocalConsumer{defaultConsumer}
}
