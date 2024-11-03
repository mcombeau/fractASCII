package main

import "math"

func mandelbrot(x, y float64, maxIter int) int {
	zR, zI := 0.0, 0.0
	for i := 0; i < maxIter; i++ {
		zR, zI = zR*zR-zI*zI+x, 2*zR*zI+y
		if zR*zR+zI*zI > 4 {
			return i
		}
	}
	return maxIter
}

func julia(x, y, cR, cI float64, maxIter int) int {
	zR, zI := x, y
	for i := 0; i < maxIter; i++ {
		zR, zI = zR*zR-zI*zI+cR, 2*zR*zI+cI
		if zR*zR+zI*zI > 4 {
			return i
		}
	}
	return maxIter
}

func burningShip(x, y float64, maxIter int) int {
	zR, zI := 0.0, 0.0
	for i := 0; i < maxIter; i++ {
		zR, zI = zR*zR-zI*zI+x, 2*abs(zR)*abs(zI)+y
		if zR*zR+zI*zI > 4 {
			return i
		}
	}
	return maxIter
}

func tricorn(x, y float64, maxIter int) int {
	zR, zI := 0.0, 0.0
	for i := 0; i < maxIter; i++ {
		zR, zI = zR*zR-zI*zI+x, -2*zR*zI+y
		if zR*zR+zI*zI > 4 {
			return i
		}
	}
	return maxIter
}

func multibrot(x, y float64, power float64, maxIter int) int {
	zR, zI := 0.0, 0.0
	for i := 0; i < maxIter; i++ {
		r := math.Pow(zR*zR+zI*zI, power/2)
		theta := math.Atan2(zI, zR) * power
		zR, zI = r*math.Cos(theta)+x, r*math.Sin(theta)+y
		if zR*zR+zI*zI > 4 {
			return i
		}
	}
	return maxIter
}

func abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}
