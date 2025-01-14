package widgets

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

type SidebarWidget struct {
	*widgets.List
}

func NewSidebarWidget() *SidebarWidget {
	self := &SidebarWidget{
		List: widgets.NewList(),
	}
	self.Title = " General Info "
	self.update()
	self.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	self.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	self.WrapText = false
	return self
}

func (self *SidebarWidget) update() {
	stat, err := nvidiasmi.GetNvidiaSmiStats()
	if err != nil {
		log.Printf("error recieved from gopsutil: %v", err)
	}

	self.Rows = []string{
		fmt.Sprintf("Driver Version: %s", stat.DriverVersion),
		fmt.Sprintf("CUDA Version: %s", stat.CudaVersion),
		fmt.Sprintf("AttachedGpus: %d", stat.AttachedGpus),
	}
}
