package consumer

type LocalMqttConsumer struct {
	MqttConsumer
}

func NewLocalMqttConsumer(localFile string, listenData func() <-chan []byte) LocalMqttConsumer {
	localConsumer := NewLocalConsumer(localFile, listenData)
	return LocalMqttConsumer{
		MqttConsumer: NewMqttConsumer(localConsumer.handleData),
	}
}
