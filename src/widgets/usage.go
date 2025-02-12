package widgets

import (
	"log"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
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
	// sinData := (func() []float64 {
	// 	n := 220
	// 	ps := make([]float64, n)
	// 	for i := range ps {
	// 		ps[i] = 1 + math.Sin(float64(i)/5)
	// 	}
	// 	return ps
	// })()
	// sinData2 := (func() []float64 {
	// 	n := 220
	// 	ps := make([]float64, n)
	// 	for i := range ps {
	// 		ps[i] = 1 + math.Sin(float64(i)/2)
	// 	}
	// 	return ps
	// })()

	self.Data = make([][]float64, 2)
	self.Data[0] = make([]float64, 10)
	self.Data[1] = make([]float64, 10)
	self.Title = title
	self.AxesColor = ui.ColorGreen
	self.LineColors[0] = ui.ColorRed
	self.BorderStyle = DEFAULT_BORDER_STYLE
	self.TitleStyle = DEFAULT_TITLE_STYLE
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

func (self *UsageWidget) update() {
	data, err := nvidiasmi.GetNvidiaSmiQueryGpu("utilization.gpu,utilization.memory")
	if err != nil {
		log.Printf("Unable to get usage data from nvidia smi query %v", err)
	} else {
		if isDataSupported(data) {
			gpuUtil, gpuMem, err := self.processData(data)
			if err != nil {
				log.Printf("Unable to parse usage data %v", err)
			} else {
				self.Data[0] = append(self.Data[0], gpuUtil)
				self.Data[1] = append(self.Data[1], gpuMem)
			}
		} else {
			log.Println("Unable to get usage data")
		}
	}
}

func (self *UsageWidget) processData(data string) (float64, float64, error) {
	parts := strings.Split(data, ",")
	gpuUtil, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, 0, err
	}
	gpuMem, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, 0, err
	}
	return gpuUtil, gpuMem, nil
}
