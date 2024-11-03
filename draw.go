package main

import (
	"flag"
	"fmt"
	"strings"
)

func drawFractal(fractalType string, maxIter int, juliaCR float64, juliaCI float64, mandelbrotPower float64) {
	var output strings.Builder
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			x := xMin + (xMax-xMin)*float64(i)/float64(width)
			y := yMin + (yMax-yMin)*float64(j)/float64(height)

			var iter int
			switch fractalType {
			case "mandelbrot", "m":
				iter = mandelbrot(x, y, mandelbrotPower, maxIter)
			case "julia", "j":
				iter = julia(x, y, juliaCR, juliaCI, maxIter)
			case "burningship", "b":
				iter = burningShip(x, y, maxIter)
			case "tricorn", "t":
				iter = tricorn(x, y, maxIter)
			default:
				fmt.Println("Unknown fractal type:", fractalType)
				flag.Usage()
				return
			}
			output.WriteRune(getIterChar(iter, maxIter))
		}
		output.WriteRune('\n')
	}

	fmt.Print("\033[H\033[2J") // Clears the terminal
	fmt.Print(output.String())
}

func getIterChar(iter, maxIter int) rune {
	if iter >= maxIter {
		iter = maxIter - 1
	}
	return rune(chars[iter*len(chars)/maxIter])
}
