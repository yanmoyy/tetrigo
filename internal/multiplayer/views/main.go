package views

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/colors"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/layout"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/widget"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type mainSelection int

const (
	msStartGame mainSelection = iota
	msExit
)

func (s mainSelection) String() string {
	switch s {
	case msStartGame:
		return "Start Game"
	case msExit:
		return "Exit"
	}
	return "Unknown"
}

var _ tea.Model = &MainModel{}
var _ Sizable = &MainModel{}

type MainModel struct {
	SizeableImpl

	selected mainSelection
	keys     *mainKeyMap
}

func NewMainModel() *MainModel {
	keys := defaultMainKeyMap()
	return &MainModel{
		selected: msStartGame,
		keys:     keys,
	}
}

// Init implements tea.Model.
func (m *MainModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(keyMsg, m.keys.Exit):
			return m, tea.Quit
		case key.Matches(keyMsg, m.keys.Up):
			return m, m.Up()
		case key.Matches(keyMsg, m.keys.Down):
			return m, m.Down()
		case key.Matches(keyMsg, m.keys.Enter):
			return m, m.Enter()
		}
	}
	return m, nil
}

// View implements tea.Model.
func (m *MainModel) View() string {
	output := lipgloss.JoinVertical(lipgloss.Center,
		titleStr,
		layout.GapV(2),
		// TODO: Delete this line once the multiplayer mode is ready for public use.
		"The Multiplayer Mode is currently under development...",
		layout.GapV(2),
		selectionButton(msStartGame.String(), m.selected == msStartGame),
		layout.GapV(1),
		selectionButton(msExit.String(), m.selected == msExit),
		layout.GapV(1),
	)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, output)
}

func selectionButton(label string, selected bool) string {
	const (
		buttonWidth  = 14
		buttonHeight = 1
	)
	textColor, borderColor := colors.White, colors.White
	if selected {
		textColor = colors.Violet
		borderColor = colors.Violet
	}
	return widget.NewButton().
		Text(label).
		TextColor(textColor).
		BorderColor(borderColor).
		Width(buttonWidth).
		Height(buttonHeight).
		View()
}

func (m *MainModel) Up() tea.Cmd {
	if m.selected == msStartGame {
		m.selected = msExit
		return nil
	}
	m.selected--
	return nil
}

func (m *MainModel) Down() tea.Cmd {
	if m.selected == msExit {
		m.selected = msStartGame
		return nil
	}
	m.selected++
	return nil
}

func (m *MainModel) Enter() tea.Cmd {
	switch m.selected {
	case msStartGame:
		return SwitchPageCmd(ViewLobby)
	case msExit:
		return tea.Quit
	}
	return nil
}
