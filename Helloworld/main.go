package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("HelloWorld-App")
	w.SetContent(widget.NewLabel("Test Data"))
	w.ShowAndRun()
}
