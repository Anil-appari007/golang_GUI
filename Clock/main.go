package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	// clock.SetText(time.Now().Format("2006-01-02 03:04:05 pm MST"))
	data := fmt.Sprintf("Time: %s", time.Now().Format("2006-01-02 03:04:05 pm MST"))
	clock.SetText(data)
}
func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("")
	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}
