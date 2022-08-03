/*
This program was created by Mark Gauda in the Summer of 2022
This is the package that can generate an image from fractal color data

Features I want to have are:
	-PNG image generation
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"time"
)

//Saves a color array as an image with a given file name and file type
func saveColorArrayAsImage(colorArray []color.Color, imageName, fileType string, width, height int) {
	img := image.NewRGBA((image.Rect(0, 0, width, height)))
	var colorArrayPosition int
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, colorArray[colorArrayPosition])
			colorArrayPosition += 1
		}
	}
	//if name is blank, give it a name
	if imageName == "" {
		imageName = "image.png"
	}
	//if suffix already attached, remove it (might be better to use regular expressions)
	if imageName[len(imageName)-4] == '.' {
		imageName = imageName[0 : len(imageName)-4]
	}

	file, err := os.Create(fmt.Sprintf(imageName + "." + fileType))

	if err != nil {
		println(err)
	}

	if strings.ToLower(fileType) == "png" {
		png.Encode(file, img)

	}
	if strings.ToLower(fileType) == "jpg" {
		jpeg.Encode(file, img, nil)
	}

}

func makeImage(imageName string) {
	if imageName == "" {
		var currentTime = time.Now()
		imageName = fmt.Sprintf("fractalImage x %.16f y %.16f %d", fractalData.centerX, fractalData.centerY, currentTime.UnixMicro())
	}
	saveColorArrayAsImage(RequestHandler(fractalData.centerX, fractalData.centerY, fractalData.zoomScale, imageWidth, imageHeight, IMAGE, MULTI_LINE, 16).imageData, imageName, "png", imageWidth, imageHeight)
}
