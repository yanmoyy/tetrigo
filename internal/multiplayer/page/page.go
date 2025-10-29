package page

import tea "github.com/charmbracelet/bubbletea"

type Page int

const (
	PageMain Page = iota
	PageLobby
	PageGame
	PageSettings
)

func (p Page) String() string {
	switch p {
	case PageMain:
		return "Main"
	case PageLobby:
		return "Lobby"
	case PageGame:
		return "Game"
	case PageSettings:
		return "Settings"
	default:
		return "Unknown"
	}
}

type SwitchPageMsg struct {
	Target Page
}

func SwitchPageCmd(target Page) tea.Cmd {
	return func() tea.Msg {
		return SwitchPageMsg{
			Target: target,
		}
	}
}

// Sizable is an interface that can be resized by its parent.
type Sizable interface {
	SetSize(width, height int)
}

// SizeableImpl is a default implementation of Sizable. It changes its width and height.
type SizeableImpl struct {
	Width  int
	Height int
}

func (s *SizeableImpl) SetSize(width, height int) {
	s.Width = width
	s.Height = height
}
