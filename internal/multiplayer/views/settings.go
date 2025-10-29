package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	tea "github.com/charmbracelet/bubbletea"
)

type SettingsModel struct {
}

// Init implements tea.Model.
func (m *SettingsModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m *SettingsModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View implements tea.Model.
func (m *SettingsModel) View() string {
	return "Settings"
}

func NewSettingsModel(_ *config.Config) (*SettingsModel, error) {
	return &SettingsModel{}, nil
}
