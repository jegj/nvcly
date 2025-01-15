package widgets

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type ListWidget struct {
	*widgets.List
}

func NewListWidget() *ListWidget {
	self := &ListWidget{
		List: widgets.NewList(),
	}
	self.Title = " Attached GPUs "
	self.update()
	self.TextStyle = ui.NewStyle(ui.ColorYellow)
	self.SelectedRow = 0 // Highlight the first row
	self.BorderStyle.Fg = ui.ColorYellow
	self.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorGreen, ui.ModifierBold)
	self.WrapText = false
	return self
}

func (self *ListWidget) update() {
	stat, err := nvidiasmi.GetNvidiaSmiStats()
	if err != nil {
		log.Printf("error recieved from gopsutil: %v", err)
	}
	for _, gpu := range stat.GPU {
		self.Rows = append(self.Rows, fmt.Sprintf("%s(%s)", gpu.ProductName, gpu.Uuid))
	}
}
