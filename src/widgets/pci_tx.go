package widgets

import (
	"fmt"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type PciTxRxWidget struct {
	*widgets.Paragraph
	updateInterval time.Duration
	tx             bool
}

func NewPciTxRxWidget(title string, updateInterval time.Duration, tx bool) *PciTxRxWidget {
	self := &PciTxRxWidget{
		Paragraph:      widgets.NewParagraph(),
		updateInterval: updateInterval,
		tx:             tx,
	}
	self.Title = title
	self.TextStyle = applyStaticDataStyle()
	self.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
	self.WrapText = false
	self.BorderStyle.Fg = ui.ColorGreen
	self.update()

	go func() {
		for range time.NewTicker(self.updateInterval).C {
			self.Lock()
			self.update()
			self.Unlock()
		}
	}()
	return self
}

func (self *PciTxRxWidget) update() {
	data, err := nvidiasmi.GetNvidiaSmiDmonQueryGpu("t", 1)
	if err != nil {
		self.TextStyle = ui.NewStyle(ui.ColorRed, ui.ColorClear, ui.ModifierBold)
		self.Text = "<err>"
	} else {
		if isDataSupported(data) {
			text := self.processData(data)
			self.Text = fmt.Sprintf("%s MB/s", text)
		} else {
			self.TextStyle = applyNoDataStyles()
			self.Text = "N/A"
		}
	}
}

func (self *PciTxRxWidget) processData(data string) string {
	parts := strings.Split(data, ",")
	partsLen := len(parts)
	if self.tx {
		return strings.TrimSpace(parts[partsLen-1])
	} else {
		return strings.TrimSpace(parts[partsLen-2])
	}
}
