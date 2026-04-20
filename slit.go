package main

import "math"

func intensity(y, y0, I0, d, l float64) float64 {
	if y == 0 {
		return 0
	} else {
		th := math.Acos(y / y0)
		return 4 * I0 * math.Cos(math.Pi*d*math.Sin(th)/l) * math.Cos(math.Pi*d*math.Sin(th)/l)
	}
}
