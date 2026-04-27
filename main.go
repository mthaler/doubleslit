package main

import (
	"gonum.org/v1/plot/plotter"
)

func main() {
	var points plotter.XYs

	for i := -100000; i <= 100000; i++ {
		x := float64(i) / 1000.0
		points = append(points, plotter.XY{
			X: x,
			Y: intensity(x, 10, 50, 1, 1, 1),
		})
	}
	b := bounds{
		xmin: -100,
		xmax: 100,
		ymin: 0,
		ymax: 0.2,
	}
	l := labels{
		x: "y",
		y: "intensity",
	}
	CreateLineplotPlot(points, "y - intensity", l, b, "intensity.png")
}
