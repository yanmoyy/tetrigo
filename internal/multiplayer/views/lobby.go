package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	tea "github.com/charmbracelet/bubbletea"
)

type LobbyModel struct {
}

// Init implements tea.Model.
func (m *LobbyModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m *LobbyModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View implements tea.Model.
func (m *LobbyModel) View() string {
	return "Lobby"
}

func NewLobbyModel(cfg *config.Config) (*LobbyModel, error) {
	return &LobbyModel{}, nil
}
