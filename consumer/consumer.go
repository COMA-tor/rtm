package consumer

type Consumer interface {
	Run() error
}

type emptyConsumer int

func (*emptyConsumer) Run() error {
	return nil
}

func NewConsumer() Consumer {
	return new(emptyConsumer)
}

type DefaultConsumer struct {
	Consumer
	listenData func() <-chan []byte
	handleData func([]byte)
}

func (d DefaultConsumer) Run() error {
	out := d.listenData()
	for {
		select {
		case data := <-out:
			d.handleData(data)
		}
	}
}

func WithDataHandler(consumer Consumer, listenData func() <-chan []byte, handleData func([]byte)) DefaultConsumer {
	return DefaultConsumer{Consumer: consumer, listenData: listenData, handleData: handleData}
}