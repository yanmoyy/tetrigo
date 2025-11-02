package views

import tea "github.com/charmbracelet/bubbletea"

type View int

const (
	ViewMain View = iota
	ViewLobby
	ViewP2P
)

func (p View) String() string {
	switch p {
	case ViewMain:
		return "Main"
	case ViewLobby:
		return "Lobby"
	case ViewP2P:
		return "P2P"
	default:
		return "Unknown"
	}
}

func GetViewModel(page View, sessionID string) tea.Model {
	var view tea.Model
	switch page {
	case ViewMain:
		view = NewMainModel()
	case ViewLobby:
		view = NewLobbyModel(sessionID)
	case ViewP2P:
		view = NewP2PModel()
	}
	return view
}

type SwitchPageMsg struct {
	Target View
}

func SwitchPageCmd(target View) tea.Cmd {
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

type SessionIDSetter interface {
	SetSessionID(string)
}
