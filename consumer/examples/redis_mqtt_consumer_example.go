package main

import (
	"fmt"
	"github.com/COMA-tor/rtm/consumer"
)

const redisHost1 = "localhost"
const redisPort1 = 6379
const clientName1 = "REDIS-CLIENT-001"

func main() {
	mqttToRedisConsumer := consumer.NewMqttToRedisConsumer(
		redisHost1,
		fmt.Sprint(redisPort1),
		clientName1,
	)

	mqttToRedisConsumer.Run()
}
