package client

type MessageHandler interface {}

type Configuration interface {}

type Client interface {
	Publish(config Configuration, payload interface{}) error
	Subscribe(config Configuration, callback MessageHandler) error
}


