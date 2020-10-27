package consumer

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/COMA-tor/rtm/data"
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
)

type AirportData struct {
	IataCode        string
	MeasurementType string
	Timestamp       int64
	Measure         float64
}

type MqttToRedisConsumer struct {
	MqttConsumer
}

func NewMqttToRedisConsumer(redisHost string, redisPort string, brokerHost string, brokerPort string, clientName string) MqttToRedisConsumer {
	return MqttToRedisConsumer{
		MqttConsumer: NewMqttConsumer(
			"/airport/#",
			brokerHost,
			brokerPort,
			newDataToRedisHandler(redisHost, redisPort, clientName),
		),
	}
}

func newDataToRedisHandler(redisHost string, redisPort string, clientName string) func(bytes []byte) {
	client := redistimeseries.NewClient(fmt.Sprintf("%s:%v", redisHost, redisPort), "nohelp", nil)

	return func(bytes []byte) {
		data := getDataFromBytes(bytes)
		keyname := strings.TrimSpace(data.MeasurementType) + ":" + data.IataCode

		createRuleIfNotExists(client, keyname)

		log.Printf("Sending to Redis... %v:%v -> %v [%v]", data.MeasurementType, data.IataCode, data.Measure, data.Timestamp)
		client.Add(keyname, data.Timestamp, data.Measure)
	}
}

func getDataFromBytes(bytes []byte) AirportData {

	dataStr := string(bytes)
	data := strings.SplitAfterN(dataStr, " ", 2)

	if len(data) != 2 {
		log.Fatal(
			errors.New("topic and payload were expected, get: " + dataStr),
		)
	}

	topic, payload := data[0], data[1]

	iataCode, measurementType := getDataFromTopic(topic)
	timestamp, measure, _ := getDataFromPayload([]byte(payload))

	return AirportData{
		IataCode:        iataCode,
		MeasurementType: measurementType,
		Timestamp:       timestamp,
		Measure:         measure,
	}
}

func getDataFromPayload(payload []byte) (int64, float64, string) {
	return data.ByteToValues(payload)
}

func getDataFromTopic(topic string) (string, string) {
	topicData := strings.Split(
		strings.Replace(
			topic, "/airport/", "", 1,
		),
		"/",
	)

	if len(topicData) != 2 {
		log.Fatal(
			errors.New("IATA code and measurement type were expected from topic, get: " + topic),
		)
	}

	return topicData[0], topicData[1]
}

func createRuleIfNotExists(client *redistimeseries.Client, keyname string) {
	_, err := client.Info(keyname)
	if err != nil {
		client.CreateKeyWithOptions(keyname, redistimeseries.DefaultCreateOptions)
		client.CreateKeyWithOptions(keyname+"_avg", redistimeseries.DefaultCreateOptions)
		client.CreateRule(keyname, redistimeseries.AvgAggregation, 60, keyname+"_avg")
	}
}
