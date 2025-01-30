package widgets

import (
	"fmt"
	"log"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type TextBoxDynamicWidget struct {
	*widgets.Paragraph
	nvidiasmiQuery   string
	updateInterval   time.Duration
	isPercentageData bool
}

func NewTextBoxDynamicWidget(title string, nvidiasmiQuery string, updateInterval time.Duration, isPercentageData bool) *TextBoxDynamicWidget {
	self := &TextBoxDynamicWidget{
		Paragraph:        widgets.NewParagraph(),
		nvidiasmiQuery:   nvidiasmiQuery,
		updateInterval:   updateInterval,
		isPercentageData: isPercentageData,
	}
	self.updateInterval = updateInterval
	self.Title = title
	self.TextStyle = STATIC_DATA_STYLE
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

func (self *TextBoxDynamicWidget) update() {
	data, err := nvidiasmi.GetNvidiaSmiQueryGpu(self.nvidiasmiQuery)
	if err != nil {
		log.Printf("error recieved from nvidiasmi query: %v", err)
	}
	if isDataSupported(data) {
		if self.isPercentageData {
			self.TextStyle = applyDataStyles(data)
			self.Text = fmt.Sprintf("%s%%", data)
		} else {
			self.TextStyle = STATIC_DATA_STYLE
			self.Text = data
		}
	} else {
		self.TextStyle = NO_DATA_STYLE
		self.Text = "N/A"
	}
}

func applyDataStyles(data string) ui.Style {
	percentage, err := strconv.Atoi(data)
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
}
