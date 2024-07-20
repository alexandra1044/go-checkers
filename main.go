package checkers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	a := app.New()
	w := a.NewWindow("Go Checkers")

	board := GenerateBoard()
	w.SetContent(board)
	w.Resize(fyne.NewSize(480, 480))

	w.ShowAndRun()
}
