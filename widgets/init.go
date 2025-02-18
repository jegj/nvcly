package widgets

import (
	"fmt"
	"log"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/jegj/nvcly/nvidiasmi"
)

var (
	ERR_QUERY_STYLE      ui.Style = ui.NewStyle(ui.ColorRed, ui.ColorClear, ui.ModifierBold)
	NO_DATA_STYLE        ui.Style = ui.NewStyle(ui.ColorMagenta, ui.ColorClear, ui.ModifierBold)
	STATIC_DATA_STYLE    ui.Style = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
	DEFAULT_BORDER_STYLE ui.Style = ui.NewStyle(ui.ColorGreen)
	DEFAULT_TITLE_STYLE  ui.Style = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
)

const (
	ERR_QUERY_TEXT string = "<err>"
	NO_DATA_TEXT   string = "N/A"
)

func isDataSupported(output string) bool {
	return !strings.Contains(output, "N/A")
}

func InitiNvcly(timeInterval time.Duration) {
	stat, err := nvidiasmi.GetNvidiaSmiStats()
	if err != nil {
		log.Fatalf("failed to initialize nvcly: %v", err)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize nvcly: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	defaultGpuName := fmt.Sprintf("%s(%s)", stat.GPU[0].ProductName, stat.GPU[0].Uuid)
	driverVersion := NewTextBox("Driver Version", stat.DriverVersion)
	cudaVersion := NewTextBox("Cuda Version", stat.CudaVersion)
	selectionMessage := NewTextBox(":Press G to switch GPUs", defaultGpuName)
	gpuUsage := NewTextBoxDynamicWidget("GPU Usage", "utilization.gpu", timeInterval, true)
	memUsage := NewTextBoxDynamicWidget("Mem Usage", "utilization.memory", timeInterval, true)
	fanSpeed := NewTextBoxDynamicWidget("Fan Speed", "fan.speed", timeInterval, true)
	performanceState := NewTextBoxDynamicWidget("Per.State", "pstate", timeInterval, false)
	pciCurrentLinkGen := NewTextBoxDynamicWidget("PCIe Curr.Link", "pcie.link.gen.current", timeInterval, false)
	pciTx := NewPciTxWidget("PCI Tx", timeInterval)
	pciRx := NewPciRxWidget("PCI Rx", timeInterval)
	processes := NewProcessesWidget("Processes", timeInterval, stat.GPU[0].Processes.ProcessInfo)
	temperature := NewTemperatureWidget("Temperature", timeInterval)
	usage := NewUsageWidget("Usage", timeInterval)

	grid.Set(
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/7, driverVersion),
			ui.NewCol(1.0/7, cudaVersion),
			ui.NewCol(5.0/7, selectionMessage),
		),
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/4, gpuUsage),
			ui.NewCol(1.0/4, memUsage),
			ui.NewCol(1.0/4, fanSpeed),
			ui.NewCol(1.0/4, temperature),
		),
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/4, performanceState),
			ui.NewCol(1.0/4, pciCurrentLinkGen),
			ui.NewCol(1.0/4, pciTx),
			ui.NewCol(1.0/4, pciRx),
		),
		ui.NewRow(2.0/4,
			ui.NewCol(1.0/1, usage),
		),
		ui.NewRow(2.0/5,
			ui.NewCol(1, processes),
		),
	)

	ticker := time.NewTicker(timeInterval).C

	ui.Render(grid)
	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			ui.Render(gpuUsage, memUsage, fanSpeed, performanceState, pciCurrentLinkGen, pciTx, pciRx, processes, temperature, usage)
		}
	}
}
