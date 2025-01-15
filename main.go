package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	nvclyw "github.com/jegj/nvcly/src/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize nvcly: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	main := nvclyw.NewListWidget()
	sidebar := nvclyw.NewSidebarWidget()

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
