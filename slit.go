package main

import "math"

func intensity(x, x0, Imax, d, l float64) float64 {
	th := math.Acos(2 * x / l)
	return Imax * math.Cos(math.Pi*d*math.Sin(th)/l) * math.Cos(math.Pi*d*math.Sin(th)/l)
}
