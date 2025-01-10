package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Create a view for the table
	if v, err := g.SetView("table", maxX/4, maxY/4, 3*maxX/4, 3*maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Simple Table"
		v.Wrap = true

		// Define table content
		headers := []string{"ID", "Name", "Age"}
		rows := [][]string{
			{"1", "Alice", "30"},
			{"2", "Bob", "25"},
			{"3", "Charlie", "35"},
		}

		// Print the table
		fmt.Fprintf(v, "%-5s %-15s %-5s\n", headers[0], headers[1], headers[2])
		fmt.Fprintln(v, "---------------------------")
		for _, row := range rows {
			fmt.Fprintf(v, "%-5s %-15s %-5s\n", row[0], row[1], row[2])
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	// Add keybinding to quit
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	// Start the GUI main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
