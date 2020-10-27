package data

import (
	"fmt"
	"strconv"
)

func ValuesToByte(tsp int64, val float64, unit string) []byte {
	return []byte(
		fmt.Sprintf(
			"%v %v %v",
			tsp,
			strconv.FormatFloat(val, 'E', -1, 32),
			unit,
		),
	)
}
