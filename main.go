package main

import (
	"fmt"

	"gonum.org/v1/plot/plotter"
)

func main() {
	var points plotter.XYs

	fmt.Println("%+v", points)
	for i := -500; i <= 500; i++ {
		x := float64(i) / 100.0
		points = append(points, plotter.XY{
			X: x,
			Y: intensity(1, 2, x, 10),
		})
	}
}
