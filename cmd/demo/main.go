// Package main contains a simple test application for bar chart
package main

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/borud/bw"
)

var (
	green  = color.RGBA{G: 0xcc, A: 0xff}
	yellow = color.RGBA{R: 0xcc, G: 0xcc, A: 0xff}
	red    = color.RGBA{R: 0xcc, A: 0xff}
)

func mapColor(v float64) color.Color {
	switch {
	case v < 0.3:
		return red
	case v < 0.6:
		return yellow
	default:
		return green
	}
}

func main() {
	app := app.New()
	win := app.NewWindow("Demo")

	numBars := 50

	bars := bw.NewBars(numBars)
	bars.Spacing = 5
	bars.BarMinSize = fyne.NewSize(15, 100)
	bars.ColorFunc = mapColor

	// update it with some data
	go func(bars *bw.Bars) {
		for f := float64(0); ; f += 0.05 {
			for i := 0; i < numBars; i++ {
				bars.SetValue(i, math.Abs(math.Sin(f+(0.1*float64(i)))))
			}
			bars.Refresh()
			time.Sleep(50 * time.Millisecond)
		}
	}(bars)

	win.SetContent(bars)
	win.ShowAndRun()
}
