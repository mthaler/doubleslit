package main

import (
	"math"
)

func intensity(y, y0, d, l, b float64) float64 {
	th := math.Acos(y / y0)
	i := math.Cos(math.Pi * d * math.Sin(th) / l)
	diff := sinc(math.Pi * b * math.Sin(th) / l)
	return i * i * diff * diff
}
