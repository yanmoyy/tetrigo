package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	tea "github.com/charmbracelet/bubbletea"
)

type GameModel struct {
}

// Init implements tea.Model.
func (m *GameModel) Init() tea.Cmd {
	panic("unimplemented")
}

// Update implements tea.Model.
func (m *GameModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (m *GameModel) View() string {
	panic("unimplemented")
}

func NewGameModel(cfg *config.Config) (*GameModel, error) {
	return &GameModel{}, nil
}
