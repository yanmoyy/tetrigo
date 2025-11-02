package views

import tea "github.com/charmbracelet/bubbletea"

var _ tea.Model = &P2PModel{}

type P2PModel struct{}

func (p *P2PModel) Init() tea.Cmd {
	return nil
}

func (p *P2PModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

func (p *P2PModel) View() string {
	return "P2P"
}

func NewP2PModel() *P2PModel {
	return &P2PModel{}
}
