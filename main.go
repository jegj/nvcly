package main

import (
	"log"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/jegj/nvcly/src/nvidiasmi"
	nvclyw "github.com/jegj/nvcly/src/widgets"
)

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

	main := nvclyw.NewGpuListWidget(stat)
	driverVersion := nvclyw.NewTextBox("Driver Version", stat.DriverVersion)
	cudaVersion := nvclyw.NewTextBox("Cuda Version", stat.CudaVersion)
	attachedGpus := nvclyw.NewTextBox("Attached Gpus", strconv.Itoa(stat.AttachedGpus))

	grid.Set(
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/3, driverVersion),
			ui.NewCol(1.0/3, cudaVersion),
			ui.NewCol(1.0/3, attachedGpus),
		),
		ui.NewRow(3.75/4,
			ui.NewCol(1, main),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for e := range uiEvents {
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(grid)
		}
	}
}
