package sensor

import (
	"github.com/aquilax/go-perlin"
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

var TemperatureGenerator GeneratorFunc = PerlinGenerator(4, 2, 3, -5, 25, 60000)
var HygrometryGenerator GeneratorFunc = PerlinGenerator(4, 2, 3, 0, 100, 60000)
var PressureGenerator GeneratorFunc = PerlinGenerator(4, 2, 3, 1000, 1024, 60000)
