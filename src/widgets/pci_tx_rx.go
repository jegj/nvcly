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
	reverseIndex   int
}

const (
	TX_REVERSE_INDEX int = 1
	RX_REVERSE_INDEX int = 2
)

func newPciTxRxWidget(title string, updateInterval time.Duration, reverseIndex int) *PciTxRxWidget {
	self := &PciTxRxWidget{
		Paragraph:      widgets.NewParagraph(),
		updateInterval: updateInterval,
		reverseIndex:   reverseIndex,
	}
	self.Title = title
	self.TextStyle = ui.NewStyle(ui.ColorCyan, ui.ColorClear, ui.ModifierBold)
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

func NewPciTxWidget(title string, updateInterval time.Duration) *PciTxRxWidget {
	return newPciTxRxWidget(title, updateInterval, TX_REVERSE_INDEX)
}

func NewPciRxWidget(title string, updateInterval time.Duration) *PciTxRxWidget {
	return newPciTxRxWidget(title, updateInterval, RX_REVERSE_INDEX)
}

func (self *PciTxRxWidget) update() {
	data, err := nvidiasmi.GetNvidiaSmiDmonQueryGpu("t", 1)
	if err != nil {
		self.TextStyle = ERR_QUERY_STYLE
		self.Text = ERR_QUERY_TEXT
	} else {
		if isDataSupported(data) {
			text := self.processData(data)
			self.Text = fmt.Sprintf("%s MB/s", text)
		} else {
			self.TextStyle = NO_DATA_STYLE
			self.Text = "N/A"
		}
	}
}

func (self *PciTxRxWidget) processData(data string) string {
	parts := strings.Split(data, ",")
	partsLen := len(parts)
	return strings.TrimSpace(parts[partsLen-self.reverseIndex])
}
