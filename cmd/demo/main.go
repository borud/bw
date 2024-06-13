// Package main contains a simple test application for bar chart
package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/borud/bw"
)

func main() {
	app := app.New()
	win := app.NewWindow("Demo")

	numBars := 13
	numCharts := 13

	vbox := container.NewVBox()

	for j := 0; j < numCharts; j++ {

		b := bw.NewBars(numBars)
		b.Spacing = 5
		b.BarMinSize = fyne.NewSize(30, 100)
		vbox.Add(b)
		go func(bars *bw.Bars) {
			for i := 0; ; i++ {
				for i := 0; i < numBars; i++ {
					bars.SetValue(i, rand.Float64())
				}
				bars.Refresh()
				time.Sleep(100 * time.Millisecond)
			}
		}(b)
	}

	win.SetContent(vbox)
	win.ShowAndRun()
}
