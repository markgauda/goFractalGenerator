/*
This program was created by Mark Gauda in the Summer of 2022
This is the package that can generate fractal time to escape matricies

Features I want to have are:
	-Time to escape matrix generator
	-Concurrentcy
*/

package main

import (
	"math/cmplx"
	"sync"
)

/*
This is the request handler. It processes a request and creates a fractal accordignly
x and y are the midpoint coordinates for the fractal
scale is how far apart the min and max values for x and y will be
width and heigth are the number of pixles to generate on the x and y axis
send determines what pipe the data will be sent to
	-GAME_WINDOW = back to the game window to be updated
	-IMAGE = to the image processor to be saved for later viewing
concurrent determines what algorithm to process the fractal with
	-NON_CONCURRENT = a non concurrent algorithm
	-MULTI_LINE = an algorithm that can generate multiple lines of the image at one time
threads determines how many threads you want to use for the generator
Return
*/
func RequestHandler(x, y, scale float64, width, heigth, send, concurrent, threads int) requestReturnData {
	var escapeMatrix []int
	var concreteRequestReturnData requestReturnData
	if concurrent == NON_CONCURRENT {
		escapeMatrix = generateFractalNotConcurrent(x, y, scale, width, heigth)
	} else if concurrent == MULTI_LINE {
		if threads < 1 {
			threads = 1
		}
		escapeMatrix = generateFractalLineConcurrent(x, y, scale, width, heigth, threads)
	}
	if send == GAME_WINDOW {
		//send escapeMatrix to encoder for game window
		concreteRequestReturnData.gameScreen = makeGameScreen(escapeMatrix)
		return concreteRequestReturnData
	} else if send == IMAGE {
		//send escapeMatrix to encoder for image
		concreteRequestReturnData.imageData = makeImageData(escapeMatrix, width, heigth)
		return concreteRequestReturnData
	}
	return concreteRequestReturnData
}

/*calcualtes how long a given coordinate will stay in the given range
Arg coordinate: The complex coordinate you are inspecting
Arg iterations: The number of iterations you are willing to try
Arg escapeSize: This is the range outside of the original coordinate you allow before considering it an escape
*/
func findTimeToEscape(coordinate complex128, iterations int, escapeSize float64) int {
	var v complex128
	for i := 0; i < iterations; i++ {
		v = v*v + coordinate
		if cmplx.Abs(v) > escapeSize {
			return i
		}
	}
	return maxIterations
}

//finds the camera bounds for a point on the graph given a zoom scale
func getMinMax(point complex128, scale float64) (float64, float64, float64, float64) {
	xMax := scale/2 + real(point)
	yMax := scale/2 + imag(point)
	xMin := real(point) - scale/2
	yMin := imag(point) - scale/2
	return xMin, xMax, yMin, yMax
}

//Generate the time to escape matrix in with a non-concurrent algorithm
func generateFractalNotConcurrent(x, y, scale float64, width, height int) []int {
	xmin, xmax, ymin, ymax := getMinMax(complex(x, y), scale)

	var escapeMatrix []int = make([]int, width*height)
	var escapeMatrixPosition int

	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			escapeMatrix[escapeMatrixPosition] = findTimeToEscape(z, maxIterations, escapeSize)
			escapeMatrixPosition++
		}
	}
	return escapeMatrix
}

//func generateFractalConcurrent

//func generateFractalLine
func generateFractalLineConcurrent(x, y, scale float64, width int, height int, threads int) []int {
	xmin, xmax, ymin, ymax := getMinMax(complex(x, y), scale)

	var escapeMatrix []int = make([]int, width*height)
	var wg sync.WaitGroup
	var workingThreads int = getWorkingThreads(height, threads)
	var linesToGo int = height - workingThreads
	var py int
	for i := 0; linesToGo >= 0; i++ {
		for j := 0; j < workingThreads; j++ {
			y := float64(py)/float64(height)*(ymax-ymin) + ymin
			wg.Add(1)
			go generateFractalLine(width, escapeMatrix[width*py:width*(py+1)], &wg, xmin, xmax, y)
			py++
		}
		wg.Wait()
		workingThreads = getWorkingThreads(linesToGo, threads)
		linesToGo -= workingThreads
	}
	return escapeMatrix
}

func getWorkingThreads(tasks, threads int) int {
	check := tasks % threads
	if check == 0 {
		return threads
	} else {
		return check
	}

}

func generateFractalLine(width int, escapeMatrix []int, wg *sync.WaitGroup, xmin, xmax, y float64) {
	defer wg.Done()
	for px := 0; px < width; px++ {
		x := float64(px)/float64(width)*(xmax-xmin) + xmin
		z := complex(x, y)
		escapeMatrix[px] = findTimeToEscape(z, maxIterations, escapeSize)
	}
}
