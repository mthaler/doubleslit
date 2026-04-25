package main

import "math"

func sinc(a float64) float64 {
	if a != 0 {
		return math.Sin(a) / a
	} else {
		return 1.0
	}
}
