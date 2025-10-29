package widget

import "github.com/charmbracelet/lipgloss"

// Button is a button that has a border and a text.

var _ Widget = &Button{}

type Button struct {
	width  int
	height int

	textColor   lipgloss.Color
	borderColor lipgloss.Color
	text        string
}

func NewButton(width, height int, textColor, borderColor lipgloss.Color, text string) *Button {
	return &Button{
		width:       width,
		height:      height,
		textColor:   textColor,
		borderColor: borderColor,
		text:        text,
	}
}

// View implements Widget.
func (b *Button) View() string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(b.borderColor).
		Foreground(b.textColor).
		Width(b.width).
		Height(b.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(b.text)
}
