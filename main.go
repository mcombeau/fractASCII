package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	width          = 80
	height         = 40
	defaultMaxIter = 50
	xMin, xMax     = -2.0, 1.0
	yMin, yMax     = -1.5, 1.5
	chars          = " .:-=+*#%@"
)

func main() {
	var output strings.Builder

	fractalType := flag.String("f", "mandelbrot", "Fractal type:\n\tm, mandelbrot (default),\n\tj, julia,\n\tb, burningship,\n\tx, multibrot\n\tt, tricorn.")
	maxIter := flag.Int("i", defaultMaxIter, "Maximum number of iterations.")
	juliaCReal := flag.Float64("jr", -0.7, "Real part of the constant for Julia set.")
	juliaCImag := flag.Float64("ji", 0.27015, "Imaginary part of the constant for Julia set.")
	multibrotPower := flag.Float64("p", 3, "Power for the Multibrot set.")
	flag.Parse()

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			x := xMin + (xMax-xMin)*float64(i)/float64(width)
			y := yMin + (yMax-yMin)*float64(j)/float64(height)

			var iter int
			switch *fractalType {
			case "mandelbrot", "m":
				iter = mandelbrot(x, y, *maxIter)
			case "julia", "j":
				iter = julia(x, y, *juliaCReal, *juliaCImag, *maxIter)
			case "burningship", "b":
				iter = burningShip(x, y, *maxIter)
			case "multibrot", "x":
				iter = multibrot(x, y, *multibrotPower, *maxIter)
			case "tricorn", "t":
				iter = tricorn(x, y, *maxIter)
			default:
				fmt.Println("Unknown fractal type:", *fractalType)
				flag.Usage()
				return
			}
			output.WriteRune(getIterChar(iter, *maxIter))
		}
		output.WriteRune('\n')
	}

	fmt.Print(output.String())
}

func getIterChar(iter, maxIter int) rune {
	if iter >= maxIter {
		iter = maxIter - 1
	}
	return rune(chars[iter*len(chars)/maxIter])
}
