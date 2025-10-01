package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	tea "github.com/charmbracelet/bubbletea"
)

type SettingsModel struct {
}

// Init implements tea.Model.
func (m *SettingsModel) Init() tea.Cmd {
	panic("unimplemented")
}

// Update implements tea.Model.
func (m *SettingsModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	panic("unimplemented")
}

// View implements tea.Model.
func (m *SettingsModel) View() string {
	panic("unimplemented")
}

func NewSettingsModel(cfg *config.Config) (*SettingsModel, error) {
	return &SettingsModel{}, nil
}
