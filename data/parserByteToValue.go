package data

import (
	"strconv"
	"strings"
)

func ByteToValues(tab []byte) (int64, float64, string) {
	var spl []string = strings.Split(string(tab), " ")
	tsp, _ := strconv.ParseInt(spl[0], 10, 64)
	value, _ := strconv.ParseFloat(spl[1], 64)
	return tsp, value, spl[2]
}
