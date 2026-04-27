package main

import (
	"gonum.org/v1/plot/plotter"
)

func main() {
	var points plotter.XYs

	for i := -10000; i <= 10000; i++ {
		x := float64(i) / 1000.0
		points = append(points, plotter.XY{
			X: x,
			Y: intensity(x, 10, 100, 0.02, 0.1, 0.1),
		})
	}
	b := bounds{
		xmin: -10,
		xmax: 10,
		ymin: 0,
		ymax: 0.2,
	}
	l := labels{
		x: "y",
		y: "intensity",
	}
	CreateLineplotPlot(points, "y - intensity", l, b, "intensity.png")
}
