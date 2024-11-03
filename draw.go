package main

import (
	"flag"
	"fmt"
	"strings"
)

func drawFractal(settings FractalSettings) {
	var output strings.Builder
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			x := xMin + (xMax-xMin)*float64(i)/float64(width)
			y := yMin + (yMax-yMin)*float64(j)/float64(height)

			var iter int
			switch settings.fractalType {
			case "mandelbrot", "m":
				settings.fractalType = "mandelbrot"
				iter = mandelbrot(x, y, settings.mandelbrotPower, settings.maxIter)
			case "julia", "j":
				settings.fractalType = "julia"
				iter = julia(x, y, settings.juliaCR, settings.juliaCI, settings.maxIter)
			case "burningship", "b":
				settings.fractalType = "burningship"
				iter = burningShip(x, y, settings.maxIter)
			case "tricorn", "t":
				settings.fractalType = "tricorn"
				iter = tricorn(x, y, settings.maxIter)
			default:
				fmt.Println("Unknown fractal type:", settings.fractalType)
				flag.Usage()
				return
			}
			output.WriteRune(getIterChar(iter, settings.maxIter))
		}
		output.WriteRune('\n')
	}

	// Clear the screen without affecting the scrollback buffer
	fmt.Print("\033[H\033[J")
	printHeader(settings)
	fmt.Print(output.String())
	printControls()
}

func printControls() {
	if hideUI {
		fmt.Printf("\n\n")
		return
	}
	fmt.Println("Controls: [k][w] up, [j][s] down, [h][a] left, [l][d] right, [+][=] zoom in, [-] zoom out")
	fmt.Println("Press [u] to hide UI")
}

func printHeader(settings FractalSettings) {
	if hideUI {
		fmt.Printf("\n\n")
		return
	}
	fmt.Printf("Fractal: %s | Iterations: %d", settings.fractalType, settings.maxIter)
	if settings.fractalType == "mandelbrot" {
		fmt.Printf(" | Mandelbrot Power: %f", settings.mandelbrotPower)
	}
	if settings.fractalType == "julia" {
		fmt.Printf(" | julia CR: %f, CI: %f", settings.juliaCR, settings.juliaCI)
	}
	fmt.Printf("\nxMin: %f, xMax: %f, yMin: %f, yMax: %f\n", xMin, xMax, yMin, yMax)
}

func getIterChar(iter, maxIter int) rune {
	if iter >= maxIter {
		iter = maxIter - 1
	}
	return rune(chars[iter*len(chars)/maxIter])
}
