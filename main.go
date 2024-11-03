package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// Rendering
	width  = 80
	height = 40
	chars  = " .:-=+*#%@"

	// Fractal math
	defaultMaxIter = 50

	// Camera
	zoomFactor = 1.1
	panFactor  = 0.1
)

var (
	// Terminal mode
	termiosBackup syscall.Termios

	// Complex mapping
	xMin, xMax = -2.0, 1.0
	yMin, yMax = -1.5, 1.5

	// UI
	hideUI = false
)

type FractalSettings struct {
	fractalType     string
	maxIter         int
	juliaCR         float64
	juliaCI         float64
	mandelbrotPower float64
}

func main() {
	settings := parseArgs()

	enableRawModeTTY()
	defer disableRawModeTTY()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		disableRawModeTTY()
		os.Exit(0)
	}()

	// Setup input handling
	fd := int(os.Stdin.Fd())
	syscall.SetNonblock(fd, true)
	buf := make([]byte, 1)

	// Initial fractal draw
	drawFractal(settings)

	for {
		// Read a single byte (key press)
		n, err := syscall.Read(fd, buf)
		if err != nil {
			if err == syscall.EAGAIN { // No input, continue without redrawing
				time.Sleep(50 * time.Millisecond)
				continue
			} else {
				fmt.Println("Error reading key:", err)
				break
			}
		}

		if n <= 0 { // No bytes read, nothing to do
			continue
		}

		// A key was pressed, redraw
		switch buf[0] {
		case 'k', 'w': // Up
			yMin -= panFactor * (yMax - yMin)
			yMax -= panFactor * (yMax - yMin)
		case 'j', 's': // Down
			yMin += panFactor * (yMax - yMin)
			yMax += panFactor * (yMax - yMin)
		case 'h', 'a': // Left
			xMin -= panFactor * (xMax - xMin)
			xMax -= panFactor * (xMax - xMin)
		case 'l', 'd': // Right
			xMin += panFactor * (xMax - xMin)
			xMax += panFactor * (xMax - xMin)
		case '+', '=': // Zoom in
			centerX := (xMin + xMax) / 2
			centerY := (yMin + yMax) / 2
			xRange := (xMax - xMin) / zoomFactor
			yRange := (yMax - yMin) / zoomFactor
			xMin, xMax = centerX-xRange/2, centerX+xRange/2
			yMin, yMax = centerY-yRange/2, centerY+yRange/2
		case '-': // Zoom out
			xRange, yRange := (xMax-xMin)*(zoomFactor-1), (yMax-yMin)*(zoomFactor-1)
			xMin, xMax = xMin-xRange, xMax+xRange
			yMin, yMax = yMin-yRange, yMax+yRange
		case 'u':
			hideUI = !hideUI
		case 'q': // Exit
			return
		default: // Invalid keypress, no need to redraw
			continue
		}

		drawFractal(settings)
	}
}

func parseArgs() (settings FractalSettings) {
	fractalType := flag.String("f", "mandelbrot", "Fractal type: [m]andelbrot, [j]ulia, [b]urningship, or [t]ricorn.")
	maxIter := flag.Int("i", defaultMaxIter, "Maximum number of iterations.")
	juliaCR := flag.Float64("jr", -0.7, "Real part of the constant for Julia set.")
	juliaCI := flag.Float64("ji", 0.27015, "Imaginary part of the constant for Julia set.")
	mandelbrotPower := flag.Float64("p", 2, "Power for the Mandelbrot ('Multibrot') set.")
	flag.Parse()

	return FractalSettings{
		fractalType:     *fractalType,
		maxIter:         *maxIter,
		juliaCR:         *juliaCR,
		juliaCI:         *juliaCI,
		mandelbrotPower: *mandelbrotPower,
	}
}
