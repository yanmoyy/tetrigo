package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
}

// Init implements tea.Model.
func (m *MainModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m *MainModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View implements tea.Model.
func (m *MainModel) View() string {
	return "main"
}

func NewMainModel(cfg *config.Config) (*MainModel, error) {
	return &MainModel{}, nil
}
