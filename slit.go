package main

import "math"

func intensity(Imax, d, th, l float64) float64 {

	return Imax * math.Cos(math.Pi*d*math.Sin(th)/l) * math.Cos(math.Pi*d*math.Sin(th)/l)
}
