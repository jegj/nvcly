package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/jegj/nvcly/src/nvidiasmi"
)

func main() {
	stat, err := nvidiasmi.GetNvidiaSmiStats()
	if err != nil {
		log.Fatalf("failed to collect nvcly stats: %v", err)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize nvcly: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	main := widgets.NewList()
	main.Title = "Attached GPUs"
	main.Rows = []string{}
	main.TextStyle = ui.NewStyle(ui.ColorYellow)
	for _, gpu := range stat.GPU {
		main.Rows = append(main.Rows, fmt.Sprintf("%s(%s)", gpu.ProductName, gpu.Uuid))
	}
	main.WrapText = false
	main.SelectedRow = 0 // Highlight the first row
	main.BorderStyle.Fg = ui.ColorYellow
	main.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorGreen, ui.ModifierBold)

	sidebar := widgets.NewList()
	sidebar.Title = "General Info"
	sidebar.Rows = []string{
		fmt.Sprintf("Driver Version: %s", stat.DriverVersion),
		fmt.Sprintf("CUDA Version: %s", stat.CudaVersion),
		fmt.Sprintf("AttachedGpus: %d", stat.AttachedGpus),
	}
	sidebar.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	sidebar.SelectedRowStyle = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	sidebar.WrapText = false

	grid.Set(
		ui.NewRow(1,
			ui.NewCol(3.0/4, main),
			ui.NewCol(1.0/4, sidebar),
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
