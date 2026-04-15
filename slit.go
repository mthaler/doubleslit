package main

import "math"

func intensity(Imax, d, x, l float64) float64 {
	th := math.Acos(2 * x / l)
	return Imax * math.Cos(math.Pi*d*math.Sin(th)/l) * math.Cos(math.Pi*d*math.Sin(th)/l)
}
