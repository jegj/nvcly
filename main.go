package main

import (
	"fmt"
	"log"
	"strconv"
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
	main := nvclyw.NewGpuListWidget(stat)
	driverVersion := nvclyw.NewTextBox("Driver Version", stat.DriverVersion)
	cudaVersion := nvclyw.NewTextBox("Cuda Version", stat.CudaVersion)
	attachedGpus := nvclyw.NewTextBox("Attached Gpus", strconv.Itoa(stat.AttachedGpus))
	selectionMessage := nvclyw.NewTextBox(":Press G to switch GPUs", defaultGpuName)
	gpuUsage := nvclyw.NewTextBoxGauge("Gpu Usage", "utilization.gpu", DEFAULT_TIME_INTERVAL)
	memUsage := nvclyw.NewTextBoxGauge("Memory Usage", "utilization.memory", DEFAULT_TIME_INTERVAL)
	encoderUsage := nvclyw.NewTextBoxGauge("Encoder Usage", "utilization.encoder", DEFAULT_TIME_INTERVAL)
	decoderUsage := nvclyw.NewTextBoxGauge("Decoder Usage", "utilization.decoder", DEFAULT_TIME_INTERVAL)
	fanSpeed := nvclyw.NewTextBoxGauge("Fan Speed", "fan.speed", DEFAULT_TIME_INTERVAL)

	grid.Set(
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/3, driverVersion),
			ui.NewCol(1.0/3, cudaVersion),
			ui.NewCol(1.0/3, attachedGpus),
		),
		ui.NewRow(0.25/4,
			ui.NewCol(1.0, selectionMessage),
		),
		ui.NewRow(0.25/4,
			ui.NewCol(1.0/7, gpuUsage),
			ui.NewCol(1.0/7, memUsage),
			ui.NewCol(1.0/7, encoderUsage),
			ui.NewCol(1.0/7, decoderUsage),
			ui.NewCol(1.0/7, fanSpeed),
		),
		ui.NewRow(3.0/4,
			ui.NewCol(1, main),
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
			ui.Render(gpuUsage, memUsage, encoderUsage, decoderUsage, fanSpeed)
		}
	}
}
