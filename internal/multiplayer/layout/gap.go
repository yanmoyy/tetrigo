package layout

import "github.com/charmbracelet/lipgloss"

// GapV (Gap Vertical) returns a string of the given height, filled with spaces.
func GapV(height int) string {
	return lipgloss.NewStyle().Height(height).Render("")
}

// GapH (Gap Horizontal) returns a string of the given width, filled with spaces.
func GapH(width int) string {
	return lipgloss.NewStyle().Width(width).Render("")
}
