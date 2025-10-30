package widget

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/colors"
	"github.com/charmbracelet/lipgloss"
)

// Button is a button that has a border and a text.

var _ Widget = &Button{}

type Button struct {
	width  int
	height int

	textColor   lipgloss.Color
	borderColor lipgloss.Color
	text        string
}

type ButtonOpt func(*Button)

func NewButton() *Button {
	return &Button{
		textColor:   colors.White, // default color
		borderColor: colors.White,
		text:        "",
	}
}

func (b *Button) Text(text string) *Button {
	b.text = text
	return b
}

func (b *Button) TextColor(color lipgloss.Color) *Button {
	b.textColor = color
	return b
}

func (b *Button) BorderColor(color lipgloss.Color) *Button {
	b.borderColor = color
	return b
}

func (b *Button) Width(width int) *Button {
	b.width = width
	return b
}

func (b *Button) Height(height int) *Button {
	b.height = height
	return b
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
