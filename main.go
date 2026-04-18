package main

import (
	"gonum.org/v1/plot/plotter"
)

func main() {
	var points plotter.XYs

	for i := -500; i <= 500; i++ {
		x := float64(i) / 100.0
		points = append(points, plotter.XY{
			X: x,
			Y: intensity(x, 0.5, 1, 5, 5e-7),
		})
	}
	b := bounds{
		xmin: 0,
		xmax: 10,
		ymin: 0,
		ymax: 10,
	}
	l := labels{
		x: "y",
		y: "intensity",
	}
	CreateLineplotPlot(points, "y - intensity", l, b, "intensity.png")
}
