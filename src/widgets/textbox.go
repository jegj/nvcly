package widgets

import (
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
	self.Title = title
	self.Text = text
	self.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	self.BorderStyle = DEFAULT_BORDER_STYLE
	self.TitleStyle = DEFAULT_TITLE_STYLE
	self.WrapText = false
	return self
}
