/*
This program was created by Mark Gauda in the Summer of 2022
This program is the second version of the go fractal generator

My intention here is to rebuild the fractal generator, so the code
not only looks nicer, but is easier to maintan and to understand

Features I want:
	-Command Line Argument control
	-Black to white fractal color generation
	-Concurrent fractal generation
	-Image output
	-Generality in the code base (break it into reusable modules)
*/

package main

import (
	"flag"
)

func main() {
	var sizeFlag = flag.Int("size", 1000, "This is the number of pixels that make the screen")
	var threadsFlag = flag.Int("threads", 1, "This is the number of threads to use when generating an image.")
	var midpointXFlag = flag.Float64("midpointX", 0, "The x value for midpoint to zoom in on, used for multi image zooms")
	var midpointYFlag = flag.Float64("midpointY", 0, "The y value for midpoint to zoom in on, used for multi image zooms")
	var maxIterationsFlag = flag.Int("maxIter", 200, "The maximum number of iterations the a point will be tested before assuming it is stable")
	var escapeSizeFlag = flag.Float64("escape", 4, "The absolute value that a point will need to exceed before considering it exploding to infinity")

	flag.Parse()

	setSize(*sizeFlag)
	setThreads(*threadsFlag)
	setXMidpoint(*midpointXFlag)
	setYMidpoint(*midpointYFlag)
	setEscapeSize(*escapeSizeFlag)
	setMaxIteration(*maxIterationsFlag)

	NewGame().Init()

}
