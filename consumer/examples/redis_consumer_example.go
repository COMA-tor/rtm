package main

import (
	"fmt"
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/aquilax/go-perlin"
	"log"
	"time"
)

type GeneratorFunc func(int64) float64

func PerlinGenerator(octaves int, alpha, beta, min, max, scaling float64) GeneratorFunc {
	seed := time.Now().Unix() + int64(min+max)
	generator := perlin.NewPerlin(alpha, beta, octaves, seed)
	return func(t int64) float64 {
		rand := (generator.Noise1D(float64(t)/scaling) + 1) / 2
		return min + rand*(max-min)
	}
}

var TemperatureGenerator GeneratorFunc = PerlinGenerator(4, 2, 3, -5, 25, 600)

const redisHost = "172.17.3.174"
const redisPort = 6379
const clientName = "REDIS-CLIENT-001"

func main() {
	// Connect to localhost with no password
	client := redistimeseries.NewClient(fmt.Sprintf("%s:%d", redisHost, redisPort), clientName, nil)
	keyname := "temperature:BDX"
	_, err := client.Info(keyname)
	if err != nil {
		client.CreateKeyWithOptions(keyname, redistimeseries.DefaultCreateOptions)
		client.CreateKeyWithOptions(keyname+"_avg", redistimeseries.DefaultCreateOptions)
		client.CreateRule(keyname, redistimeseries.AvgAggregation, 60, keyname+"_avg")
	}

	for {
		timestamp := time.Now().Unix()
		temp := TemperatureGenerator(timestamp)
		client.AddAutoTs(keyname, temp)

		log.Printf("Sent to Redis : %v [%v]", temp, timestamp)

		time.Sleep(time.Second)
	}
}
