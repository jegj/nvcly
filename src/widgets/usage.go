package widgets

import (
	"math"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type UsageWidget struct {
	*widgets.Plot
	updateInterval time.Duration
}

func NewUsageWidget(title string, updateInterval time.Duration) *UsageWidget {
	self := &UsageWidget{
		Plot:           widgets.NewPlot(),
		updateInterval: updateInterval,
	}
	sinData := (func() []float64 {
		n := 220
		ps := make([]float64, n)
		for i := range ps {
			ps[i] = 1 + math.Sin(float64(i)/5)
		}
		return ps
	})()

	self.Title = title
	self.Data = make([][]float64, 1)
	self.Data[0] = sinData
	self.AxesColor = ui.ColorGreen
	self.LineColors[0] = ui.ColorRed
	self.Marker = widgets.MarkerDot
	self.BorderStyle = DEFAULT_BORDER_STYLE
	self.TitleStyle = DEFAULT_TITLE_STYLE

	return self
}
