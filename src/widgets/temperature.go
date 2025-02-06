package widgets

import (
	"fmt"
	"time"

	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type TemperatureWidget struct {
	*widgets.Paragraph
	updateInterval time.Duration
}

func NewTemperatureWidget(title string, updateInterval time.Duration) *TemperatureWidget {
	self := &TemperatureWidget{
		Paragraph:      widgets.NewParagraph(),
		updateInterval: updateInterval,
	}
	self.updateInterval = updateInterval
	self.Title = title

	go func() {
		for range time.NewTicker(self.updateInterval).C {
			self.Lock()
			self.update()
			self.Unlock()
		}
	}()
	return self
}

func (self *TemperatureWidget) update() {
	data, err := nvidiasmi.GetNvidiaSmiQueryGpu("temperature.gpu")
	if err != nil {
		self.TextStyle = ERR_QUERY_STYLE
		self.Text = ERR_QUERY_TEXT
	} else {
		if isDataSupported(data) {
			self.TextStyle = STATIC_DATA_STYLE
			self.Text = fmt.Sprintf("%s °C", data)
		} else {
			self.TextStyle = NO_DATA_STYLE
			self.Text = NO_DATA_TEXT
		}
	}
}
