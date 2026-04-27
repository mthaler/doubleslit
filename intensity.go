package main

import (
	"math"
)

func intensity(y, y0, x0, d, l, b float64) float64 {
	th := math.Atan(y / x0)
	si := math.Sin(th)
	//i := math.Cos(math.Pi * d * math.Sin(th) / l)
	//diff := sinc(math.Pi * b * math.Sin(th) / l)
	//return i * i * diff * diff
	return math.Pi * d * si
}
