package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GenerateBoard() *fyne.Container {

	grid := container.NewGridWithColumns(8)

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			canvas.NewRectangle(color.Black)
			//background rectangle
			bg := canvas.NewRectangle(color.Black)
			if x%2 == y%2 {
				bg.FillColor = color.White
			}
			counter := canvas.NewImageFromResource(GetCounterSVG(Standard, Black))
			//maintain aspect ratio of image when stretching the canvas
			counter.FillMode = canvas.ImageFillContain
			grid.Add(container.NewStack(bg, counter))

		}
	}
	return grid
}
