package widgets

import (
	"log"
	"strconv"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/nvidiasmi"
)

type UsageWidget struct {
	*widgets.Plot
	updateInterval time.Duration
	counter        int32
}

func NewUsageWidget(title string, updateInterval time.Duration) *UsageWidget {
	u := &UsageWidget{
		Plot:           widgets.NewPlot(),
		updateInterval: updateInterval,
		counter:        0,
	}

	u.Data = make([][]float64, 1)
	u.Data[0] = make([]float64, 250)
	u.Title = title
	u.AxesColor = ui.ColorGreen
	u.LineColors[0] = ui.ColorRed
	u.BorderStyle = DEFAULT_BORDER_STYLE
	u.TitleStyle = DEFAULT_TITLE_STYLE
	u.update()
	go func() {
		for range time.NewTicker(u.updateInterval).C {
			u.Lock()
			u.update()
			u.Unlock()
		}
	}()

	return u
}

func (u *UsageWidget) update() {
	data, err := nvidiasmi.GetNvidiaSmiQueryGpu("utilization.gpu,utilization.memory")
	if err != nil {
		log.Printf("Unable to get usage data from nvidia smi query %v", err)
	} else {
		if isDataSupported(data) {
			gpuUtil, _, err := u.processData(data)
			if err != nil {
				log.Printf("Unable to parse usage data %v", err)
			} else {
				u.Data[0][u.counter%250] = gpuUtil
				u.counter++
			}
		} else {
			log.Println("Unable to get usage data")
		}
	}
}

func (u *UsageWidget) processData(data string) (float64, float64, error) {
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
