// Package main contains a simple test application for bar chart
package main

import (
	"image/color"
	"math"
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
		b.BarMinSize = fyne.NewSize(30, 50)
		b.ColorFunc = mapColor

		vbox.Add(b)
		go func(bars *bw.Bars) {
			for f := float64(0); ; f += 0.1 {
				for i := 0; i < numBars; i++ {
					bars.SetValue(i, math.Abs(math.Sin(f+(0.1*float64(i)))))
				}
				bars.Refresh()
				time.Sleep(50 * time.Millisecond)
			}
		}(b)
	}

	win.SetContent(vbox)
	win.ShowAndRun()
}

var (
	green  = color.RGBA{G: 0xcc, A: 0xff}
	yellow = color.RGBA{R: 0xcc, G: 0xcc, A: 0xff}
	red    = color.RGBA{R: 0xcc, A: 0xff}
)

func mapColor(v float64) color.Color {
	switch {
	case v < 0.2:
		return red
	case v < 0.4:
		return yellow
	default:
		return green
	}
}
