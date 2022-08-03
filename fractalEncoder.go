/*
This program was created by Mark Gauda in the Summer of 2022
This is the program that can encode data from a time to escape matrix

Features I want to have are:
	-Black to white color generation
	-Byte matrix generator
*/

package main

import "image/color"

type byteColor struct {
	red   byte
	green byte
	blue  byte
}

/*
var (
	blackAndWhitePaletteGameWindow [maxIterations]byte
)

func init() {
	for i := range blackAndWhitePaletteGameWindow {
		blackAndWhitePaletteGameWindow[i] = (0xff - byte(float64(i)/float64(len(blackAndWhitePaletteGameWindow))*0xff))
	}
}
*/
/*
This will generate a color fro the game screen
The game screen uses hexadecimal RGB color values
*/
func makeGameScreenColor(it int) (r, g, b byte) {
	if it == maxIterations {
		return 0x00, 0x00, 0x00
	}
	//c := blackAndWhitePaletteGameWindow[it]
	c := (0xff - byte(float64(it)/float64(contrast)*0xff))
	return c, c, c
}

func makeGameScreen(escapeMatrix []int) []byteColor {
	var byteColorMatrix []byteColor = make([]byteColor, width*height)
	for i := range escapeMatrix {
		var red, green, blue = makeGameScreenColor(escapeMatrix[i])
		byteColorMatrix[i].red = red
		byteColorMatrix[i].green = green
		byteColorMatrix[i].blue = blue

	}
	return byteColorMatrix
}

func makeImageColor(it int) color.Color {
	if it == maxIterations {
		return color.Black
	}
	return color.Gray{uint8(255 - int(contrast)*it)}
}

func makeImageData(escapeMatrix []int, width, height int) []color.Color {
	var imageData []color.Color = make([]color.Color, width*height)
	for i := range escapeMatrix {
		imageData[i] = makeImageColor(escapeMatrix[i])
	}
	return imageData
}
