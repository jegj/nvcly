package widgets

import (
	"strings"

	ui "github.com/gizak/termui/v3"
)

var (
	ERR_QUERY_STYLE   ui.Style = ui.NewStyle(ui.ColorRed, ui.ColorClear, ui.ModifierBold)
	NO_DATA_STYLE     ui.Style = ui.NewStyle(ui.ColorMagenta, ui.ColorClear, ui.ModifierBold)
	STATIC_DATA_STYLE ui.Style = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
)

const (
	ERR_QUERY_TEXT string = "<err>"
	NO_DATA_TEXT   string = "N/A"
)

func isDataSupported(output string) bool {
	return !strings.Contains(output, "N/A")
}
