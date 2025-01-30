package widgets

import (
	"strings"

	ui "github.com/gizak/termui/v3"
)

var (
	NO_DATA_STYLE     ui.Style = ui.NewStyle(ui.ColorMagenta, ui.ColorClear, ui.ModifierBold)
	STATIC_DATA_STYLE ui.Style = ui.NewStyle(ui.ColorYellow, ui.ColorClear, ui.ModifierBold)
)

func isDataSupported(output string) bool {
	return !strings.Contains(output, "N/A")
}
