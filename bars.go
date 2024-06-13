// Package bars contains a simple bargraph widget
package bw

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Bars implements a simple bar chart widget
type Bars struct {
	widget.BaseWidget
	numBars    int
	mu         sync.RWMutex
	values     []float64
	BarMinSize fyne.Size
	Spacing    float32
}

type barsRenderer struct {
	bars    *Bars
	barFgs  []*canvas.Rectangle
	barBgs  []*canvas.Rectangle
	objects []fyne.CanvasObject
}

const (
	defaultBarWidth   = 20.0
	defaultBarHeight  = 100.0
	defaultBarSpacing = 10.0
)

// NewBars creates new bars widget
func NewBars(n int) *Bars {
	bars := &Bars{
		numBars: n,
		values:  make([]float64, n),
		BarMinSize: fyne.Size{
			Width:  defaultBarWidth,
			Height: defaultBarHeight,
		},
		Spacing: defaultBarSpacing,
	}
	bars.ExtendBaseWidget(bars)
	return bars
}

// SetValue sets the value at a given index i to value v.
func (b *Bars) SetValue(i int, v float64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.values[i] = v
}

// Value returns the value at a given index.
func (b *Bars) Value(i int) float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.values[i]
}

// CreateRenderer returns a renderer for Bars
func (b *Bars) CreateRenderer() fyne.WidgetRenderer {
	r := &barsRenderer{
		bars:   b,
		barFgs: make([]*canvas.Rectangle, b.numBars),
		barBgs: make([]*canvas.Rectangle, b.numBars),
	}

	for i := 0; i < b.numBars; i++ {
		r.barBgs[i] = canvas.NewRectangle(theme.InputBackgroundColor())
		r.barFgs[i] = canvas.NewRectangle(theme.PrimaryColor())

		r.objects = append(r.objects, r.barBgs[i], r.barFgs[i])
	}
	return r
}

func (r *barsRenderer) Layout(size fyne.Size) {
	for i := 0; i < r.bars.numBars; i++ {

		vsplit := size.Height * float32(r.bars.Value(i))

		barWidth := ((size.Width / float32(r.bars.numBars)) - r.bars.Spacing) + (r.bars.Spacing / float32(r.bars.numBars))

		// Set the sizes
		r.barBgs[i].Resize(fyne.NewSize(barWidth, size.Height-vsplit))
		r.barFgs[i].Resize(fyne.NewSize(barWidth, vsplit))

		// Position
		xpos := (barWidth + r.bars.Spacing) * float32(i)
		r.barBgs[i].Move(fyne.NewPos(xpos, 0.0))
		r.barFgs[i].Move(fyne.NewPos(xpos, size.Height-vsplit))
	}
}

func (r *barsRenderer) Refresh() {
	r.Layout(r.bars.Size())
}

func (r *barsRenderer) MinSize() fyne.Size {
	width := (r.bars.BarMinSize.Width * float32(r.bars.numBars)) + (r.bars.Spacing * float32(r.bars.numBars-1))
	return fyne.NewSize(width, r.bars.BarMinSize.Height)
}

func (r *barsRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *barsRenderer) Destroy() {
	// does nothing
}
