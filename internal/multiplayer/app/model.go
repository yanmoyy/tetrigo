package app

import (
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/views"
	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = &model{}

type model struct {
	app   *App
	id    string
	child tea.Model

	width  int
	height int

	ExitError error
}

func initialModel() *model {
	m := &model{}
	m.child = views.NewMainModel()
	return m
}

func (m model) Init() tea.Cmd {
	return m.initChild()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ErrMsg:
		m.ExitError = msg
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case views.SwitchPageMsg:
		m.child = views.GetViewModel(msg.Target, m.id)
		cmd := m.child.Init()
		return m, cmd

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if sizable, ok := m.child.(views.Sizable); ok {
			sizable.SetSize(m.width, m.height)
		}

	case ProgMsg:
		m.app.send(msg.ID, msg.Msg)
	}

	var cmd tea.Cmd
	m.child, cmd = m.child.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.child.View()
}

func (m *model) initChild() tea.Cmd {
	var cmds []tea.Cmd
	cmd := m.child.Init()
	cmds = append(cmds, cmd)
	if sizable, ok := m.child.(views.Sizable); ok {
		sizable.SetSize(m.width, m.height)
	}
	m.child, cmd = m.child.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

type ErrMsg error
type ProgMsg struct {
	ID  progID
	Msg tea.Msg
}
