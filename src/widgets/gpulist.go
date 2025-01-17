package widgets

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type GpuListWidget struct {
	*widgets.List
}

func NewGpuListWidget(stat *nvidiasmi.NvidiaSmiLog) *GpuListWidget {
	self := &GpuListWidget{
		List: widgets.NewList(),
	}
	self.Title = " Attached GPUs "
	self.update(stat)
	self.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	self.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
	self.SelectedRow = 0 // Highlight the first row
	self.BorderStyle.Fg = ui.ColorGreen
	self.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorGreen, ui.ModifierBold)
	self.WrapText = false
	return self
}

func (self *GpuListWidget) update(stat *nvidiasmi.NvidiaSmiLog) {
	for _, gpu := range stat.GPU {
		self.Rows = append(self.Rows, fmt.Sprintf("%s(%s)", gpu.ProductName, gpu.Uuid))
	}
}
