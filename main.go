package main

import (
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/jegj/nvcly/src/nvidiasmi"
	nvclyw "github.com/jegj/nvcly/src/widgets"
)

const DEFAULT_TIME_INTERVAL time.Duration = time.Second

func main() {
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
	driverVersion := nvclyw.NewTextBox("Driver Version", stat.DriverVersion)
	cudaVersion := nvclyw.NewTextBox("Cuda Version", stat.CudaVersion)
	selectionMessage := nvclyw.NewTextBox(":Press G to switch GPUs", defaultGpuName)
	gpuUsage := nvclyw.NewTextBoxDynamicWidget("GPU Usage", "utilization.gpu", DEFAULT_TIME_INTERVAL, true)
	memUsage := nvclyw.NewTextBoxDynamicWidget("Mem Usage", "utilization.memory", DEFAULT_TIME_INTERVAL, true)
	fanSpeed := nvclyw.NewTextBoxDynamicWidget("Fan Speed", "fan.speed", DEFAULT_TIME_INTERVAL, true)
	performanceState := nvclyw.NewTextBoxDynamicWidget("Per.State", "pstate", DEFAULT_TIME_INTERVAL, false)
	pciCurrentLinkGen := nvclyw.NewTextBoxDynamicWidget("PCIe Curr.Link", "pcie.link.gen.current", DEFAULT_TIME_INTERVAL, false)
	pciTx := nvclyw.NewPciTxRxWidget("PCI Tx", DEFAULT_TIME_INTERVAL, true)
	pciRx := nvclyw.NewPciTxRxWidget("PCI Rx", DEFAULT_TIME_INTERVAL, false)
	processes := nvclyw.NewProcessesWidget("Processes", DEFAULT_TIME_INTERVAL, stat.GPU[0].Processes.ProcessInfo)

	grid.Set(
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/7, driverVersion),
			ui.NewCol(1.0/7, cudaVersion),
			ui.NewCol(5.0/7, selectionMessage),
		),
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/7, gpuUsage),
			ui.NewCol(1.0/7, memUsage),
			ui.NewCol(1.0/7, fanSpeed),
			ui.NewCol(1.0/7, performanceState),
			ui.NewCol(1.0/7, pciCurrentLinkGen),
			ui.NewCol(1.0/7, pciTx),
			ui.NewCol(1.0/7, pciRx),
		),
		ui.NewRow(3.0/4,
			ui.NewCol(1, processes),
		),
	)

	ticker := time.NewTicker(DEFAULT_TIME_INTERVAL).C

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
			ui.Render(gpuUsage, memUsage, fanSpeed, performanceState, pciCurrentLinkGen, pciTx, pciRx, processes)
		}
	}
}
