/*
This program was created by Mark Gauda in the Summer of 2022
This is the package that can generate fractal time to escape matricies

Features I want to have are:
	-Time to escape matrix generator
	-Concurrentcy

Known bugs
	-It appears that the arbitrary precision generator does not pass
	the y value to the concurrent line generator function reliably.
	Adding a dely of 5ms between the update of the y value and the
	passing of the y value makes the bug much less noticable. I'm
	guessing this is an issue with my implimentation of concurrency
	-Some of the lines are out of place in the final image, I'm guessing
	there is some logic error that gives the line generator the wrong y
	coordinate, because whole lines on the y axis are out of order.
	-Prining the value of some big.Float numbers causes a panic I have
	not yet determined what about these pareticular numbers causes the panic
	but the panic has always been due to an index out of bounds error
*/

package main

import (
	"fmt"
	"math/big"
	"math/cmplx"
	"sync"
	"time"
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
		if !arbitraryPrecision {
			escapeMatrix = generateFractalLineConcurrent(x, y, scale, width, heigth, threads)
		} else {
			escapeMatrix = generateFractalLineConcurrentArbitraryPrecision(*big.NewFloat(x), *big.NewFloat(y), scale, width, heigth, threads)
		}
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

func findTimeToEscapeArbitraryPrecision(coordinate arbPrecComplex, iterations int, escapeSize big.Float) int {
	var v arbPrecComplex
	var vAbs big.Float
	for i := 0; i < iterations; i++ {
		v = v.multiply(v)
		v = v.add(coordinate)
		vAbs = v.abs()
		if vAbs.Cmp(&escapeSize) == 1 { //if v < escapseSize
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

func getMinMaxArbitraryPrecision(point arbPrecComplex, scale float64) (big.Float, big.Float, big.Float, big.Float) {
	var xMax big.Float
	var xMin big.Float
	var yMax big.Float
	var yMin big.Float

	xMax.Add(big.NewFloat(scale/2), &point.real)
	yMax.Add(big.NewFloat(scale/2), &point.imaginary)
	xMin.Sub(&point.real, big.NewFloat(scale/2))
	yMin.Sub(&point.imaginary, big.NewFloat(scale/2))
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

func generateFractalLineConcurrentArbitraryPrecision(x, y big.Float, scale float64, width int, height int, threads int) []int {
	var midpoint arbPrecComplex = arbPrecComplex{x, y}
	var yMaxMinusyMin, pyOverheight, intermediateMultiplication big.Float
	xmin, xmax, ymin, ymax := getMinMaxArbitraryPrecision(midpoint, scale)
	yMaxMinusyMin = *yMaxMinusyMin.Sub(&ymax, &ymin)

	var escapeMatrix []int = make([]int, width*height)
	var wg sync.WaitGroup
	var workingThreads int = getWorkingThreads(height, threads)
	var linesToGo int = height - workingThreads
	var py int
	for i := 0; linesToGo >= 0; i++ {
		for j := 0; j < workingThreads; j++ {
			//y := float64(py)/float64(height)*(ymax-ymin) + ymin
			//yMaxMinusyMin = *yMaxMinusyMin.Sub(&ymax, &ymin)
			pyOverheight.Set(big.NewFloat(float64(py) / float64(height)))
			intermediateMultiplication = *intermediateMultiplication.Mul(&pyOverheight, &yMaxMinusyMin)
			y = *y.Add(&intermediateMultiplication, &ymin)
			fmt.Printf("out, %s ", y.Text('f', 10)) //debug
			time.Sleep(time.Millisecond * 5)        //debug
			wg.Add(1)
			go generateFractalLineArbitraryPrecision(width, escapeMatrix[width*py:width*(py+1)], &wg, xmin, xmax, y)
			py++
		}
		fmt.Printf("\n") //debug
		wg.Wait()
		workingThreads = getWorkingThreads(linesToGo, threads)
		linesToGo -= workingThreads
	}
	return escapeMatrix
}

func generateFractalLineArbitraryPrecision(width int, escapeMatrix []int, wg *sync.WaitGroup, xmin, xmax, y big.Float) {
	defer wg.Done()
	//fmt.Printf("in, %s ", y.Text('f', 10)) //debug
	var xMaxMinusxMin big.Float
	for px := 0; px < width; px++ {
		var x big.Float
		xMaxMinusxMin.Sub(&xmax, &xmin)
		x.Mul(big.NewFloat(float64(px)/float64(width)), &xMaxMinusxMin)
		x.Add(&x, &xmin) // := float64(px)/float64(width)*(xmax-xmin) + xmin
		z := arbPrecComplex{x, y}
		escapeMatrix[px] = findTimeToEscapeArbitraryPrecision(z, maxIterations, *big.NewFloat(escapeSize))
	}
}
