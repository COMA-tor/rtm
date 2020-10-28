package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Airport struct {
	IataCode string `json:"iata_code"`
}

var ctx = context.Background()

func main() {
	jsonFile, err := os.Open("../airport-codes_json.json")

	if err != nil {
		log.Println(err)
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	keys := make([]Airport, 0)
	if err := decoder.Decode(&keys); err == io.EOF {
		return
	} else if err != nil {
		log.Fatal(err)
	}

	filteredResult := Filter(keys, func(airport Airport) bool {
		return airport.IataCode != ""
	})

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	setResults := make(map[string]interface{}, 0)
	for index, airport := range filteredResult {
		setResults[strconv.Itoa(index)] = airport.IataCode
	}

	err = rdb.HSet(ctx, "iata", setResults).Err()
	if err != nil {
		panic(err)
	}

}

func Filter(vs []Airport, f func(Airport) bool) []Airport {
	vsf := make([]Airport, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
