package widgets

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type TextBoxGaugeWidget struct {
	*widgets.Paragraph
	updateInterval time.Duration
}

func NewTextBoxGauge(title string, text string, updateInterval time.Duration) *TextBoxGaugeWidget {
	self := &TextBoxGaugeWidget{
		Paragraph: widgets.NewParagraph(),
	}
	self.updateInterval = updateInterval
	self.Title = fmt.Sprintf("  %s  ", title)
	self.Text = text
	self.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	self.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
	self.WrapText = false
	self.BorderStyle.Fg = ui.ColorGreen

	go func() {
		for range time.NewTicker(self.updateInterval).C {
			self.Lock()
			self.update()
			self.Unlock()
		}
	}()
	return self
}

func (self *TextBoxGaugeWidget) update() {
	utilization, err := nvidiasmi.GetNvidiaSmiQueryGpu()
	if err != nil {
		log.Printf("error recieved from nvidiasmi query: %v", err)
	}
	self.Text = utilization
}
