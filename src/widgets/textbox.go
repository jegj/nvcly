package widgets

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type TextBoxWidget struct {
	*widgets.Paragraph
}

func NewTextBox(title string, text string) *TextBoxWidget {
	self := &TextBoxWidget{
		Paragraph: widgets.NewParagraph(),
	}
	self.Title = fmt.Sprintf("  %s  ", title)
	self.Text = text
	self.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	self.TitleStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)
	self.WrapText = false
	self.BorderStyle.Fg = ui.ColorGreen
	return self
}
