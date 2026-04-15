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
			Y: intensity(1, 2, x, 10),
		})
	}
	b := bounds{
		xmin: 0,
		xmax: 10,
		ymin: 0,
		ymax: 10,
	}
	l := labels{
		x: "x",
		y: "intensity",
	}
	CreateLineplotPlot(points, "x - intensity", l, b, "intensity.png")
}
