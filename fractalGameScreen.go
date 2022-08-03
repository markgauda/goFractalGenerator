/*
This program was created by Mark Gauda in the Summer of 2022
This program will handle game related events such as
the window, and input controls

This module was inspiered by the ebiten mandelbrot example
*/

package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//var screenWidth, screenHeight int = width, height
var tickPerSecondCap = 30 //number of times the game updates

func (FractalMetaData) calculateYCoordinate(pixleCoord int) float64 {
	_, _, yMax, yMin := getMinMax(complex(fractalData.centerX, fractalData.centerY), fractalData.zoomScale)
	//yMax := minMaxArray[1]
	//yMin := minMaxArray[3]
	yRange := yMax - yMin
	coordinate := float64(pixleCoord)/float64(height)*yRange + yMin
	return coordinate
}

func (FractalMetaData) calculateXCoordinate(pixleCoord int) float64 {
	xMin, xMax, _, _ := getMinMax(complex(fractalData.centerX, fractalData.centerY), fractalData.zoomScale)
	//yMax := minMaxArray[1]
	//yMin := minMaxArray[3]
	xRange := xMax - xMin
	coordinate := float64(pixleCoord)/float64(height)*xRange + xMin
	return coordinate
}

type Game struct {
	offscreen    *ebiten.Image
	offscreenPix []byte
	keys         []ebiten.Key
}

type FractalMetaData struct {
	centerX   float64
	centerY   float64
	zoomScale float64
}

//This makes a new game and returns it
func NewGame() *Game {
	game := &Game{
		offscreen:    ebiten.NewImage(width, height),
		offscreenPix: make([]byte, width*height*4),
	}
	game.updateOffScreen()
	return game
}

func SetScreenWidth(screenSize int) {
	if screenSize > 0 {
		height = screenSize
		width = screenSize
		ebiten.SetWindowSize(width, height)
	}
}

func (game *Game) updateOffScreen() {
	var screenColor []byteColor = RequestHandler(fractalData.centerX, fractalData.centerY, fractalData.zoomScale, width, height, GAME_WINDOW, MULTI_LINE, threads).gameScreen
	for i, color := range screenColor {
		/*
			We multiply the index by 4 because we do not want to
			overwrite the previous iteration data, and because
			the color information needs to be in an array of bytes
			byte i*4+0 = pixel[i].red
			byte i*4+1 = pixel[i].green
			byte i*4+2 = pixel[i].blue
			byte i*4+3 = pixel[i].alpha
		*/
		game.offscreenPix[i*4] = color.red
		game.offscreenPix[i*4+1] = color.green
		game.offscreenPix[i*4+2] = color.blue
		game.offscreenPix[i*4+3] = 0xff

	}
	game.offscreen.ReplacePixels(game.offscreenPix)
}

//This is needed for the Game interface
func (game *Game) Update() error {
	game.keys = inpututil.AppendPressedKeys(game.keys[:0])

	for _, key := range game.keys {
		if key == ebiten.KeyZ {
			fractalData.zoomScale *= 0.95
		} else if key == ebiten.KeyX {
			fractalData.zoomScale *= (1 / 0.95)
		} else if key == ebiten.KeyArrowUp {
			fractalData.centerY = fractalData.centerY - (1 * 0.025 * fractalData.zoomScale)
		} else if key == ebiten.KeyArrowDown {
			fractalData.centerY = fractalData.centerY + (1 * 0.025 * fractalData.zoomScale)
		} else if key == ebiten.KeyArrowLeft {
			fractalData.centerX = fractalData.centerX - (1 * 0.025 * fractalData.zoomScale)
		} else if key == ebiten.KeyArrowRight {
			fractalData.centerX = fractalData.centerX + (1 * 0.025 * fractalData.zoomScale)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			decreaseMaxIteration()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
			increaseMaxIteration()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			contrast *= 2
		} else if inpututil.IsKeyJustPressed(ebiten.KeyO) {
			if contrast/2 > 0 {
				contrast /= 2
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			makeImage("")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			imageScale += 1
			changeImageSize()
		} else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			if imageScale-1 > 0 {
				imageScale -= 1
				changeImageSize()
			}
		}
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.updateOffScreen()
	cursorX, cusrsorY := ebiten.CursorPosition()
	ebitenutil.DebugPrint(game.offscreen, fmt.Sprintf(
		"fractal location(x=%.16f, y=%.16f, scale=%.16f)\ncursor position(x=%.16f, y=%.16f)\niterations=%d, contrast=%f, screen shot multiplier=%d\nMove = arrow keys, zoomIn/zoomOut = z/x, screenShot = spaceBar\nincrease/decrease screenshot size = t/r, increase/decrease iteration = w/q\nincrease/decrease contrast = p/o",
		fractalData.centerX, -fractalData.centerY, fractalData.zoomScale, fractalData.calculateXCoordinate(cursorX), fractalData.calculateYCoordinate(cusrsorY), maxIterations, contrast, imageScale))
	screen.DrawImage(game.offscreen, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (game *Game) Init() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Mandelbrot (Demo)")
	ebiten.SetMaxTPS(tickPerSecondCap)
	//ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMinimum) //only update window when needed
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
