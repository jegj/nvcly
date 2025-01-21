package widgets

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type TextBoxGaugeWidget struct {
	*widgets.Paragraph
	nvidiasmiQuery string
	updateInterval time.Duration
}

func NewTextBoxGauge(title string, nvidiasmiQuery string, updateInterval time.Duration) *TextBoxGaugeWidget {
	self := &TextBoxGaugeWidget{
		Paragraph:      widgets.NewParagraph(),
		nvidiasmiQuery: nvidiasmiQuery,
		updateInterval: updateInterval,
	}
	self.updateInterval = updateInterval
	self.Title = fmt.Sprintf("  %s  ", title)
	self.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
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

func (self *TextBoxGaugeWidget) update() {
	utilization, err := nvidiasmi.GetNvidiaSmiQueryGpu(self.nvidiasmiQuery)
	if err != nil {
		log.Printf("error recieved from nvidiasmi query: %v", err)
	}
	self.TextStyle = applyStyles(utilization)
	self.Text = fmt.Sprintf("%s%%", utilization)
}

func applyStyles(strPercentage string) ui.Style {
	if isDataSupported(strPercentage) {
		percentage, err := strconv.Atoi(strPercentage)
		if err != nil {
			log.Printf("error casting nvidia-smi stat into number: %v", err)
		}

		if percentage >= 0 && percentage < 35 {
			return ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
		} else if percentage >= 35 && percentage < 70 {
			return ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
		} else {
			return ui.NewStyle(ui.ColorRed, ui.ColorClear, ui.ModifierBold)
		}
	} else {
		return ui.NewStyle(ui.ColorMagenta, ui.ColorClear, ui.ModifierBold)
	}
}

func isDataSupported(output string) bool {
	return !strings.Contains(output, "N/A")
}
