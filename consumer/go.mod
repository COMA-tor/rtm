module github.com/COMA-tor/rtm/consumer

go 1.15

replace github.com/COMA-tor/rtm/mqttclient => ../mqttclient

require (
	github.com/COMA-tor/rtm/mqttclient v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201002202402-0a1ea396d57c // indirect
)
