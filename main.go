package main

import (
	"gonum.org/v1/plot/plotter"
)

func main() {
	var points plotter.XYs

	for i := -1000; i <= 1000; i++ {
		x := float64(i) / 1000.0
		points = append(points, plotter.XY{
			X: x,
			Y: intensity(x, 10, 50, 1, 1, 1),
		})
	}
	b := bounds{
		xmin: -1,
		xmax: 1,
		ymin: 0,
		ymax: 1,
	}
	l := labels{
		x: "y",
		y: "intensity",
	}
	CreateLineplotPlot(points, "y - intensity", l, b, "intensity.png")
}
