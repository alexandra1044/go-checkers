package main

import (
	"go-checkers/checkers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	w := a.NewWindow("Go Checkers")

	board := checkers.GenerateBoard()
	w.SetContent(board)
	w.Resize(fyne.NewSize(480, 480))

	w.ShowAndRun()
}
