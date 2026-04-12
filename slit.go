package main

import "math"

func intensity(I0, a float64) float64 {
	return 4 * I0 * math.Cos(a/2.0) * math.Cos(a/2.0)
}
