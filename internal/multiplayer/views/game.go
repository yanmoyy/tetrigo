package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	tea "github.com/charmbracelet/bubbletea"
)

type GameModel struct {
}

// Init implements tea.Model.
func (m *GameModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m *GameModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View implements tea.Model.
func (m *GameModel) View() string {
	return "Game"
}

func NewGameModel(_ *config.Config) (*GameModel, error) {
	return &GameModel{}, nil
}
