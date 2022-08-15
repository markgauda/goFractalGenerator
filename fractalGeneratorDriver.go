/*
This program was created by Mark Gauda in the Summer of 2022
This is the file where all the package level variables are declared
and where all the usable functions are revieled

*/

package main

import "image/color"

var width, height int
var threads int
var contrast float64 = 8
var imageScale = 2
var imageWidth, imageHeight int
var arbitraryPrecision bool = false

type requestReturnData struct {
	gameScreen []byteColor
	imageData  []color.Color
}

const (
	GAME_WINDOW = iota + 1
	IMAGE
)

const (
	NON_CONCURRENT = iota + 1
	MULTI_LINE
)

var fractalData FractalMetaData = FractalMetaData{
	centerX:   -1.2567441202584817,
	centerY:   -0.3854166876371618,
	zoomScale: 4,
}

var maxIterations int = 200

var escapeSize float64 = 4

var showDebugInfo = false

//func MakeFractalImage()

func setSize(size int) {
	width = size
	height = size
	changeImageSize()
}

func setThreads(numOfThreads int) {
	threads = numOfThreads
}

func setXMidpoint(newXMid float64) {
	fractalData.centerX = newXMid
}

func setYMidpoint(newYMid float64) {
	fractalData.centerY = newYMid
}

func setEscapeSize(newEscapeSize float64) {
	escapeSize = newEscapeSize
}

func setMaxIteration(newMaxIteration int) {
	maxIterations = newMaxIteration
}

func increaseMaxIteration() {
	maxIterations *= 2
}

func decreaseMaxIteration() {
	if maxIterations/2 > 0 {
		maxIterations /= 2
	}
}

func changeImageSize() {
	imageWidth, imageHeight = width*imageScale, height*imageScale
}
