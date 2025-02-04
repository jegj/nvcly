package widgets

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type ProcessesWidget struct {
	*widgets.Table
	updateInterval time.Duration
}

func NewProcessesWidget(title string, updateInterval time.Duration, initialData []nvidiasmi.NvidiaSmiProcessInfo) *ProcessesWidget {
	self := &ProcessesWidget{
		Table:          widgets.NewTable(),
		updateInterval: updateInterval,
	}
	self.updateInterval = updateInterval
	self.Title = title
	self.RowSeparator = true
	self.BorderStyle = ui.NewStyle(ui.ColorGreen)
	self.generateRows(initialData)
	self.TextStyle = STATIC_DATA_STYLE
	self.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
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

func (self *ProcessesWidget) update() {
	stat, err := nvidiasmi.GetNvidiaSmiStats()
	if err != nil {
		log.Fatalf("failed to initialize nvcly: %v", err)
	} else {
		self.generateRows(stat.GPU[0].Processes.ProcessInfo)
	}
}

func (self *ProcessesWidget) generateRows(initialData []nvidiasmi.NvidiaSmiProcessInfo) {
	self.Rows = [][]string{}
	self.Rows = append(self.Rows, []string{"Pid", "Type", "ProcessName", "UsedMemory"})
	for _, processInfo := range initialData {
		self.Rows = append(self.Rows, []string{processInfo.Pid, processInfo.Type, processInfo.ProcessName, processInfo.UsedMemory})
	}
}
