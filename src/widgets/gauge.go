package widgets

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type GaugeWidget struct {
	*widgets.Gauge
}

func NewGauge(title string) *GaugeWidget {
	self := &GaugeWidget{
		Gauge: widgets.NewGauge(),
	}
	self.Title = fmt.Sprintf("  %s  ", title)
	self.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)

	self.Percent = 84
	/*
		self.BarColor = ui.ColorYellow
		self.LabelStyle = ui.NewStyle(ui.ColorBlue)
	*/

	self.BorderStyle.Fg = ui.ColorGreen
	self.BarColor = ui.ColorRed

	return self
}
