package bw

import (
	"fmt"
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ColorFunc takes a value v in the range [0.0,1.0] and returns a color. You can
// use this to render the color of the bar according to some scheme of your own
// choosing.
type ColorFunc func(v float64) color.Color

// Bars implements a simple bar chart widget
type Bars struct {
	widget.BaseWidget
	numBars        int
	mu             sync.RWMutex
	values         []float64
	BarMinSize     fyne.Size
	Spacing        float32
	ColorFunc      ColorFunc
	ShowPercentage bool
}

type barsRenderer struct {
	bars           *Bars
	percentageText []*canvas.Text
	barFgs         []*canvas.Rectangle
	barBgs         []*canvas.Rectangle
	objects        []fyne.CanvasObject
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
		bars:           b,
		percentageText: make([]*canvas.Text, b.numBars),
		barFgs:         make([]*canvas.Rectangle, b.numBars),
		barBgs:         make([]*canvas.Rectangle, b.numBars),
	}

	for i := 0; i < b.numBars; i++ {
		r.percentageText[i] = canvas.NewText("", theme.ForegroundColor())
		r.percentageText[i].Alignment = fyne.TextAlignCenter

		r.barBgs[i] = canvas.NewRectangle(theme.InputBackgroundColor())
		r.barFgs[i] = canvas.NewRectangle(theme.PrimaryColor())

		r.objects = append(r.objects, r.barBgs[i], r.barFgs[i], r.percentageText[i])
	}
	return r
}

func (r *barsRenderer) Layout(size fyne.Size) {
	for i := 0; i < r.bars.numBars; i++ {
		value := r.bars.Value(i)

		vsplit := size.Height * float32(value)

		barWidth := ((size.Width / float32(r.bars.numBars)) - r.bars.Spacing) + (r.bars.Spacing / float32(r.bars.numBars))

		// Set the sizes
		r.barBgs[i].Resize(fyne.NewSize(barWidth, size.Height-vsplit))
		r.barFgs[i].Resize(fyne.NewSize(barWidth, vsplit))

		// Position
		xpos := (barWidth + r.bars.Spacing) * float32(i)
		r.barBgs[i].Move(fyne.NewPos(xpos, 0.0))
		r.barFgs[i].Move(fyne.NewPos(xpos, size.Height-vsplit))

		if r.bars.ShowPercentage {
			r.percentageText[i].Show()
			r.percentageText[i].Text = fmt.Sprintf("%3d%%", int(value*100))
			r.percentageText[i].Resize(fyne.NewSize(barWidth, 0))
			r.percentageText[i].TextSize = barWidth * 0.3
			r.percentageText[i].Move(fyne.NewPos(xpos, (size.Height/2)-theme.Padding()))
			r.percentageText[i].Refresh()
		} else {
			r.percentageText[i].Hide()
		}

		// Set custom color
		if r.bars.ColorFunc != nil {
			r.barFgs[i].FillColor = r.bars.ColorFunc(value)
		}
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
