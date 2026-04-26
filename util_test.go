package main

import (
	"math"
	"testing"
)

var sincTests = []struct {
	in       float64
	expected float64
}{
	{0.0, 1.0}, {math.Pi, 0.0},
}

func TestSinc(t *testing.T) {
	for _, s := range sincTests {
		t.Run("sinc", func(t *testing.T) {
			actual := sinc(s.in)
			if actual < 0.0001 {
				actual = math.Round(actual)
			}
			if actual != s.expected {
				t.Errorf("Wrong result for sinc(%g): expected %g, got %g", s.in, s.expected, actual)
			}
		})
	}
}
