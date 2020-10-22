package client

// Configuration of the client used to publish and subscribe.
type Configuration interface{}

// Payload that will be published by the client.
type Payload interface{}

// MessageHandler define what to do when a message is received.
type MessageHandler interface{}

// Client is the interface definition for a client as used by this library.
type Client interface {
	// Publish will publish a message with the specified configuration.
	// Returns an error if the client was not able to deliver the message.
	Publish(config Configuration, payload Payload) error
	// Subscribe starts a new subscription. Provide a message handler to execute when
	// a message is published.
	Subscribe(config Configuration, callback MessageHandler) error
}
