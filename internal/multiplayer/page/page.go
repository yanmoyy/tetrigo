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
